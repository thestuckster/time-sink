import {Box, Paper, Tab, Tabs, Typography} from "@mui/material";
import React, {useEffect, useState} from "react";
import {bindings} from "../../wailsjs/go/models";
import {GetUsageBetweenDates} from "../../wailsjs/go/bindings/UsageBinding";
import {ResponsiveContainer,} from "recharts";
import UsageBar from "./UsageBar";
import {convertDataSecondsToHours} from "../utils/timeUtils";
import UsageInfo = bindings.UsageInfo;

const data = [
  {
    name: "Page A",
    uv: 4000,
    pv: 2400,
    amt: 2400,
  },
  {
    name: "Page B",
    uv: 3000,
    pv: 1398,
    amt: 2210,
  },
  {
    name: "Page C",
    uv: 2000,
    pv: 9800,
    amt: 2290,
  },
  {
    name: "Page D",
    uv: 2780,
    pv: 3908,
    amt: 2000,
  },
  {
    name: "Page E",
    uv: 1890,
    pv: 4800,
    amt: 2181,
  },
  {
    name: "Page F",
    uv: 2390,
    pv: 3800,
    amt: 2500,
  },
  {
    name: "Page G",
    uv: 3490,
    pv: 4300,
    amt: 2100,
  },
];

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


  useEffect(() => {
    fetchDailyUsage();
    setInterval(fetchDailyUsage, 90000);

    fetchLast30DayUsage();
    setInterval(fetchLast30DayUsage, 90000);

    fetchLastYear();
    setInterval(fetchLastYear, 90000);
  }, [])


  const onTabChange = (event: React.SyntheticEvent, tab: number): void => {
    setTabValue(tab);
  }


  return (
    <div style={{textAlign: "center", justifyContent: "center",}}>
      <Paper elevation={4}>
        <Box>
          <Tabs value={tabValue} onChange={onTabChange}>
            <Tab label={"Today's Usage"} {...tabProps(0)} />
            <Tab label={"Last 30 Days"}  {...tabProps(1)} />
            <Tab label={"Yearly Usage"}  {...tabProps(2)} />
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
        </Box>
      </Paper>

    </div>
  );
}