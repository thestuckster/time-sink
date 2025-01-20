import {Box, Paper, Tab, Table, TableContainer, Tabs, Typography} from "@mui/material";
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import React, {useEffect, useState} from "react";
import {bindings} from "../../wailsjs/go/models";
import {GetAllTimeUsage, GetUsageBetweenDates} from "../../wailsjs/go/bindings/UsageBinding";
import {ResponsiveContainer,} from "recharts";
import UsageBar from "./UsageBar";
import {convertDataSecondsToHours} from "../utils/timeUtils";
import UsageInfo = bindings.UsageInfo;

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

function CustomTabPanel(props: TabPanelProps) {
  const {children, value, index, ...other} = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && <Box sx={{p: 3}}>{children}</Box>}
    </div>
  );
}

function tabProps(index: number) {
  return {
    id: `simple-tab-${index}`,
    'aria-controls': `simple-tabpanel-${index}`,
  };
}

export default function StatsContent() {
  const [tabValue, setTabValue] = useState<number>(0);
  const [dailyUsage, setDailyUsage] = useState<UsageInfo[]>([]);
  const [last30, setLast30,] = useState<UsageInfo[]>([]);
  const [lastYear, setLastYear] = useState<UsageInfo[]>([]);
  const [totalAllTime, setTotalAllTime] = useState<UsageInfo[]>();

  const fetchDailyUsage = () => {
    console.info("fetchDailyUsage");
    const now = new Date();
    now.setHours(0, 0, 0, 0);
    let tomorrow = new Date(now)
    tomorrow.setDate(now.getDate() + 1);

    GetUsageBetweenDates(now, tomorrow).then(usageInfo => {
      setDailyUsage(convertDataSecondsToHours(usageInfo));
    });
  }

  const fetchLast30DayUsage = () => {
    console.info("fetchLast30DayUsage");
    const now = new Date();
    now.setHours(0, 0, 0, 0);
    now.setDate(now.getDate() + 1);

    const then = new Date()
    then.setHours(0, 0, 0, 0);
    then.setDate(now.getDate() - 31);

    GetUsageBetweenDates(then, now).then(usageInfo => {
      setLast30(convertDataSecondsToHours(usageInfo));
    });
  }

  const fetchLastYear = () => {
    console.info("fetchLastYear");
    const now = new Date();
    now.setHours(0, 0, 0, 0);
    now.setDate(now.getDate() + 1);

    const then = new Date()
    then.setHours(0, 0, 0, 0);
    then.setDate(now.getDate() - 365);

    GetUsageBetweenDates(then, now).then(usageInfo => {
      setLastYear(convertDataSecondsToHours(usageInfo));
    });
  }

  const getAllTimeUsage = () => {
    console.info("getAllTimeUsage");
    GetAllTimeUsage().then(usageInfo => {
      setTotalAllTime(convertDataSecondsToHours(usageInfo));
    });
  }

  useEffect(() => {
    fetchDailyUsage();
    setInterval(fetchDailyUsage, 90000);

    fetchLast30DayUsage();
    setInterval(fetchLast30DayUsage, 90000);

    fetchLastYear();
    setInterval(fetchLastYear, 90000);

    getAllTimeUsage();
    setInterval(getAllTimeUsage, 90000);
  }, [])

  const onTabChange = (event: React.SyntheticEvent, tab: number): void => {
    setTabValue(tab);
  }

  const buildAllTimeTable = () =>{
    if(totalAllTime) {
      return totalAllTime.map((row) => (
        <TableRow
          key={row.Name}
          // sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
        >
          <TableCell component="th" scope="row">
            {row.Name}
          </TableCell>
          <TableCell>{row.Duration}</TableCell>
        </TableRow>
      ))
    }
  }

  return (
    <div style={{textAlign: "center", justifyContent: "center",}}>
      <Paper elevation={4}>
        <Box>
          <Tabs value={tabValue} onChange={onTabChange}>
            <Tab label={"Today's Usage"} {...tabProps(0)} />
            <Tab label={"Last 30 Days"}  {...tabProps(1)} />
            <Tab label={"Yearly Usage"}  {...tabProps(2)} />
            <Tab label={"All Time"}  {...tabProps(3)} />
          </Tabs>
          <CustomTabPanel
            index={0}
            value={tabValue}
          >
            <Typography variant="h5">
              Today's Usage In Hours
            </Typography>
            <ResponsiveContainer width={"100%"} style={{margin: "auto"}}>
              <UsageBar data={dailyUsage}/>
            </ResponsiveContainer>
          </CustomTabPanel>
          <CustomTabPanel index={1} value={tabValue}>
            <Typography variant="h5">
              Last 30 Days of Usage In Hours
            </Typography>
            <UsageBar data={last30}/>
          </CustomTabPanel>
          <CustomTabPanel index={2} value={tabValue}>
            <Typography variant="h5">
              Usage Over The Year
            </Typography>
            <UsageBar data={lastYear}/>
          </CustomTabPanel>
          <CustomTabPanel index={3} value={tabValue}>
            <Typography variant="h5">
              Total Usage
            </Typography>
            <TableContainer component={Paper}>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>Name</TableCell>
                    <TableCell>Total Running Time (Hours) </TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {buildAllTimeTable()}
                </TableBody>
              </Table>
            </TableContainer>
          </CustomTabPanel>
        </Box>
      </Paper>
    </div>
  );
}