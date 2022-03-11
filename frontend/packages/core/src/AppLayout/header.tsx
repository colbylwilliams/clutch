import React from "react";
import { Link } from "react-router-dom";
import styled from "@emotion/styled";
import { AppBar as MuiAppBar, Box, Grid, Toolbar, Typography } from "@material-ui/core";

import { useStorageContext } from "../Contexts";
import { Alert } from "../Feedback";
import { FeatureOn, SimpleFeatureFlag } from "../flags";
import { NPSHeader } from "../NPS";

import Logo from "./logo";
import Notifications from "./notifications";
import SearchField from "./search";
import ShortLinker from "./shortLinker";
import { UserInformation } from "./user";

export const APP_BAR_HEIGHT = "64px";

const AppBar = styled(MuiAppBar)({
  minWidth: "fit-content",
  background: "linear-gradient(90deg, #38106b 4.58%, #131c5f 89.31%)",
  zIndex: 1201,
  height: APP_BAR_HEIGHT,
});

// Since the AppBar is fixed we need a div to take up its height in order to push other content down.
const ClearAppBar = styled.div({ height: APP_BAR_HEIGHT });

const Title = styled(Typography)({
  margin: "12px 0px 12px 8px",
  fontWeight: "bold",
  fontSize: "30px",
  paddingLeft: "5px",
  color: "rgba(255, 255, 255, 0.87)",
});

const Header: React.FC = () => {
  const showNotifications = false;
  const { shortLinked } = useStorageContext();

  return (
    <>
      <AppBar position="fixed" elevation={0}>
        <Toolbar>
          <Link to="/">
            <Logo />
          </Link>
          <Title>clutch</Title>
          <Grid container alignItems="center" justify="flex-end">
            {shortLinked && (
              <Alert severity="warning" collapsible hover>
                You have visited this page from a ShortLink, your local changes will not be saved
                while here. Please reload to continue as normal.
              </Alert>
            )}
            <Box>
              <SearchField />
            </Box>
            <SimpleFeatureFlag feature="shortLinks">
              <FeatureOn>
                <ShortLinker />
              </FeatureOn>
            </SimpleFeatureFlag>
            {showNotifications && <Notifications />}
            <SimpleFeatureFlag feature="npsHeader">
              <FeatureOn>
                <NPSHeader />
              </FeatureOn>
            </SimpleFeatureFlag>
            <UserInformation />
          </Grid>
        </Toolbar>
      </AppBar>
      <ClearAppBar />
    </>
  );
};

export default Header;
