import React from 'react';
import {AppBar, Button, IconButton, Toolbar, Typography} from "@mui/material";
import MenuIcon from "@mui/icons-material/Menu";
import StatsContent from "./components/StatsContent";
import ConfigContent from "./components/ConfigContent";
import UsageBar from "./components/UsageBar";

function App() {

  const [currentSelection, setCurrentSelection] = React.useState<string>("stats");

  const onStatsClick = () => {
    setCurrentSelection("stats");
  }

  const onConfigClick= () => {
    setCurrentSelection("config");
  }

  const renderContent = () => {
    if (currentSelection === "stats") {
      return <StatsContent />;
    }

    return <ConfigContent />;
  }

  return (
    <>
      <AppBar position="static">
        <Toolbar disableGutters>
          <Typography variant="h6">Time Sink</Typography>
          <Button color="inherit" onClick={onStatsClick}>Stats</Button>
          <Button color="inherit" onClick={onConfigClick}>Config</Button>
        </Toolbar>
      </AppBar>
      <br />
      {renderContent()}
      {/*<UsageBar />*/}
    </>
  );
}

export default App
