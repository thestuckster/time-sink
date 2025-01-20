import {
  Accordion, AccordionActions,
  AccordionDetails,
  AccordionSummary, Box, Button,
  Card,
  CardContent,
  CardHeader, Modal,
  Paper, Table, TableContainer, TextField,
  Typography
} from "@mui/material";
import React, {useEffect, useState} from "react";
import {GetConfig, SaveConfig} from "../../wailsjs/go/bindings/TimeSinkConfigBinding";
import {bindings} from "../../wailsjs/go/models";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import ConfigDto = bindings.ConfigDto;
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";
import TableBody from "@mui/material/TableBody";
import {GetCurrentlyRunningProcesses} from "../../wailsjs/go/bindings/ProcessBinding";
import ProcessListDto = bindings.ProcessListDto;
import ProcessDto = bindings.ProcessDto;

export default function ConfigContent() {

  const [config, setConfig] = useState<ConfigDto>();
  const [updateInterval, setUpdateInterval] = useState<string>();
  const [opened, setOpened] = useState<boolean>(false);
  const [runningProcesses, setRunningProcesses] = useState<ProcessDto[]>([]);

  useEffect(() => {
      GetConfig().then(config => {
        setConfig(config);
      })
  }, []);

  useEffect(() => {
    GetCurrentlyRunningProcesses().then(processes => {
      setRunningProcesses(processes.processes);
    })
  }, [opened])

  const onUpdateIntervalChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    console.log("interval changed ", e.currentTarget.value);
    setUpdateInterval(e.currentTarget.value);
  }

  const onSave = () => {
    console.log("on save");
    config!.check_interval = updateInterval!;
    SaveConfig(config!).then(() => console.info("Config Updated"));
  }

  const buildTrackedApplicationsTable = () => {
    //TODO: ideally, I'd like to add a checkbox to each row and a button to delete the checked rows from the config when
    // a "Stop Tracking" button is pressed.
    if(config) {
      return config.applications.map((appName) => (
        <TableRow
          key={appName}
          // sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
        >
          <TableCell component="th" scope="row">
            {appName}
          </TableCell>
        </TableRow>
      ))
    }
  }

  const toggleApplicationTrackingModal = () => {
    setOpened(!opened);
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
            <Button variant="contained" onClick={onSave}>Save</Button>
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
            <TableContainer component={Paper}>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>Name</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {buildTrackedApplicationsTable()}
                </TableBody>
              </Table>
            </TableContainer>
          </AccordionDetails>
          <AccordionActions>
            <Button variant="contained" onClick={toggleApplicationTrackingModal}>Track New Application</Button>
          </AccordionActions>
        </Accordion>
      </Paper>
      <Modal
        open={opened}
        onClose={toggleApplicationTrackingModal}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Paper>
          {
            //TODO: turn this into a table with check box functionality mentioned in the other TODO but instead of removing, we are adding to the config
            runningProcesses.map((p) => (
             <p>{p.name}</p>
            ))
          }
        </Paper>
      </Modal>
    </>
  );
}