import React from "react";
import { BrowserRouter as Router, Outlet, Route, Routes } from "react-router-dom";
import _ from "lodash";

import AppLayout from "../AppLayout";
import { StorageContext } from "../Contexts";
import { ApplicationContext } from "../Contexts/app-context";
import {
  removeLocalData,
  retrieveData,
  retrieveLocalData,
  storeLocalData,
} from "../Contexts/storage-context/helpers";
import storageContextReducer from "../Contexts/storage-context/reducer";
import type { StorageContextProps } from "../Contexts/storage-context/types";
import { defaultStorageState } from "../Contexts/storage-context/types";
import { FEATURE_FLAG_POLL_RATE, featureFlags } from "../flags";
import Landing from "../landing";
import NotFound from "../not-found";

import { registeredWorkflows } from "./registrar";
import { Theme } from "./themes";
import type { ConfiguredRoute, Workflow, WorkflowConfiguration } from "./workflow";
import ErrorBoundary from "./workflow";

export interface UserConfiguration {
  [packageName: string]: {
    [key: string]: ConfiguredRoute;
  };
}

/**
 * Filter workflow routes using available feature flags.
 * @param workflows a list of valid Workflow objects.
 */
const featureFlagFilter = (workflows: Workflow[]): Promise<Workflow[]> => {
  return featureFlags().then(flags =>
    workflows.filter(workflow => {
      /* eslint-disable-next-line no-param-reassign */
      workflow.routes = workflow.routes.filter(route => {
        const show =
          route.featureFlag === undefined ||
          (flags?.[route.featureFlag] !== undefined &&
            flags[route.featureFlag].booleanValue === true);
        return show;
      });
      return workflow.routes.length !== 0;
    })
  );
};

interface ClutchAppProps {
  availableWorkflows: {
    [key: string]: () => WorkflowConfiguration;
  };
  configuration: UserConfiguration;
}

const ClutchApp: React.FC<ClutchAppProps> = ({
  availableWorkflows,
  configuration: userConfiguration,
}) => {
  const [workflows, setWorkflows] = React.useState<Workflow[]>([]);
  const [isLoading, setIsLoading] = React.useState<boolean>(true);

  const [storageState, dispatch] = React.useReducer(storageContextReducer, defaultStorageState);

  const storageProviderProps: StorageContextProps = {
    shortLinked: storageState.shortLinked,
    storeData: (componentName: string, key: string, data: any, localStorage?: boolean) =>
      dispatch({ type: "STORE_DATA", payload: { componentName, key, data, localStorage } }),
    storeLocalData,
    removeData: (componentName: string, key: string, localStorage?: boolean) =>
      dispatch({ type: "REMOVE_DATA", payload: { componentName, key, localStorage } }),
    removeLocalData,
    retrieveData: (...args) => retrieveData(storageState, ...args),
    retrieveLocalData,
    clearShortLink: (route: string) => dispatch({ type: "CLEAR_SHORT_LINK", payload: { route } }),
    clearTempData: () => dispatch({ type: "EMPTY_TEMP_DATA" }),
    tempData: () => defaultStorageState.tempStore,
  };

  const loadWorkflows = () => {
    registeredWorkflows(availableWorkflows, userConfiguration, [featureFlagFilter]).then(w => {
      setWorkflows(w);
      setIsLoading(false);
    });
  };

  React.useEffect(() => {
    loadWorkflows();
    const interval = setInterval(loadWorkflows, FEATURE_FLAG_POLL_RATE);
    return () => clearInterval(interval);
  }, []);

  const [discoverableWorkflows, setDiscoverableWorkflows] = React.useState([]);
  React.useEffect(() => {
    /** Filter out all of the workflows that are configured to be `hideNav: true`.
     * This prevents the workflows from being discoverable by the user from the UI,
     * both search and drawer navigation.
     *
     * The routes for all configured workflows will still be reachable
     * by manually providing the full path in the URI.
     */
    const pw = _.cloneDeep(workflows).filter(workflow => {
      const publicRoutes = workflow.routes.filter(route => {
        return !(route?.hideNav !== undefined ? route.hideNav : false);
      });
      workflow.routes = publicRoutes; /* eslint-disable-line no-param-reassign */
      return publicRoutes.length !== 0;
    });
    setDiscoverableWorkflows(pw);
  }, [workflows]);

  return (
    <Router>
      <Theme variant="light">
        <div id="App">
          <ApplicationContext.Provider value={{ workflows: discoverableWorkflows }}>
            <StorageContext.Provider value={storageProviderProps}>
              <Routes>
                <Route path="/" element={<AppLayout isLoading={isLoading} />}>
                  <Route key="landing" path="" element={<Landing />} />
                  {workflows.map((workflow: Workflow) => {
                    const workflowPath = workflow.path.replace(/^\/+/, "").replace(/\/+$/, "");
                    const workflowKey = workflow.path.split("/")[0];
                    return (
                      <Route
                        path={`${workflowPath}/`}
                        key={workflowKey}
                        element={
                          <ErrorBoundary workflow={workflow}>
                            <Outlet />
                          </ErrorBoundary>
                        }
                      >
                        {workflow.routes.map(route => {
                          const heading = route.displayName
                            ? `${workflow.displayName}: ${route.displayName}`
                            : workflow.displayName;
                          return (
                            <Route
                              key={workflow.path}
                              path={`${route.path.replace(/^\/+/, "").replace(/\/+$/, "")}`}
                              element={React.cloneElement(<route.component />, {
                                ...route.componentProps,
                                heading,
                              })}
                            />
                          );
                        })}
                        <Route key={`${workflow.path}/notFound`} path="*" element={<NotFound />} />
                      </Route>
                    );
                  })}
                  <Route key="notFound" path="*" element={<NotFound />} />
                </Route>
              </Routes>
            </StorageContext.Provider>
          </ApplicationContext.Provider>
        </div>
      </Theme>
    </Router>
  );
};

export default ClutchApp;
