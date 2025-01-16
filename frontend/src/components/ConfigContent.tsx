import {
  Accordion, AccordionActions,
  AccordionDetails,
  AccordionSummary, Button,
  Card,
  CardContent,
  CardHeader,
  Paper, TextField,
  Typography
} from "@mui/material";
import React, {useEffect, useState} from "react";
import {GetConfig, SaveConfig} from "../../wailsjs/go/bindings/TimeSinkConfigBinding";
import {bindings} from "../../wailsjs/go/models";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import ConfigDto = bindings.ConfigDto;

export default function ConfigContent() {

  const [config, setConfig] = useState<ConfigDto>();
  const [updateInterval, setUpdateInterval] = useState<string>();

  useEffect(() => {
      GetConfig().then(config => {
        setConfig(config);
      })
  }, []);

  const onUpdateIntervalChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    console.log("interval changed ", e.currentTarget.value);
    setUpdateInterval(e.currentTarget.value);
  }

  const onSave = () => {
    console.log("on save");
    config!.check_interval = updateInterval!;
    SaveConfig(config!).then(() => console.info("Config Updated"));
  }

  return (
    <>
      <Paper elevation={4} style={{textAlign: "center", justifyContent: "center"}}>
        <Typography variant="h6">Edit Config</Typography>
        <Accordion defaultExpanded>
          <AccordionSummary
            expandIcon={<ExpandMoreIcon />}
            aria-controls="panel1-content"
            id="panel1-header"
          >
            <Typography component="span">Update Interval</Typography>
          </AccordionSummary>
          <AccordionDetails>
            {config?.check_interval}
            <TextField
              id="updateIntervalField"
              required
              label={"Interval"}
              defaultValue={config?.check_interval}
              helperText={"An integer followed by a time unit. Ex: '1 m' to update every minute. Supported values are s (seconds), m (minutes), h (hours)."}
              onChange={onUpdateIntervalChange}
            />
          </AccordionDetails>
          <AccordionActions>
            <Button onClick={onSave}>Save</Button>
          </AccordionActions>
        </Accordion>
        <Accordion defaultExpanded>
          <AccordionSummary
            expandIcon={<ExpandMoreIcon />}
            aria-controls="panel1-content"
            id="panel1-header"
          >
            <Typography component="span">Applications</Typography>
          </AccordionSummary>
          <AccordionDetails>
            {config?.check_interval}
          </AccordionDetails>
        </Accordion>
      </Paper>
    </>
  );
}