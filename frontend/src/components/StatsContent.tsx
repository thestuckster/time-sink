import {Box, Card, Paper, Tab, Tabs, Typography} from "@mui/material";
import React, {useEffect, useState} from "react";
import {bindings} from "../../wailsjs/go/models";
import ProcessUsageData = bindings.ProcessUsageData;
import UsageInfo = bindings.UsageInfo;
import {GetUsageBetweenDates} from "../../wailsjs/go/bindings/UsageBinding";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  Rectangle,
  ResponsiveContainer,
} from "recharts";
import UsageBar from "./UsageBar";

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
  const [last30, setLast30, ] = useState<UsageInfo[]>([]);

  useEffect(() => {
    const now = new Date();
    now.setHours(0,0,0,0);
    let tomorrow = new Date(now)
    tomorrow.setDate(now.getDate() + 1);

    GetUsageBetweenDates(now, tomorrow).then(data => {
      console.info(data);
      setDailyUsage(data);
    });

    const sooner = new Date(now);
    sooner.setDate(now.getDate() - 30);

    console.log("looking with dates ", sooner, new Date());
    GetUsageBetweenDates(sooner, new Date()).then(data => {
      console.log("last 30 days", data);
      setLast30(data);
    });
  }, [])

  const onTabChange = (event: React.SyntheticEvent, tab: number): void => {
    setTabValue(tab);
  }

  const updateInfo = () => {
    console.log("Updating Info...");
    const now = new Date();
    now.setHours(0,0,0,0);
    let tomorrow = new Date(now)
    tomorrow.setDate(now.getDate() + 1);

    GetUsageBetweenDates(now, tomorrow).then(data => {
      console.info(data);
      setDailyUsage(data);
    });
  }

  setInterval(updateInfo, 90000);

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
              Today's Usage
            </Typography>
            <ResponsiveContainer width={"100%"} style={{margin: "auto"}}>
              <UsageBar data={dailyUsage}/>
            </ResponsiveContainer>
          </CustomTabPanel>
          <CustomTabPanel index={1} value={tabValue}>
            <Typography variant="h5">
              Today's Usage
            </Typography>
            <UsageBar data={last30}/>
          </CustomTabPanel>
          {/*<CustomTabPanel index={2} value={tabValue}>*/}
          {/*  <Card>*/}
          {/*    <BarChart*/}
          {/*      xAxis={[{ scaleType: 'band', data: ['group A', 'group B', 'group C'] }]}*/}
          {/*      series={[{ data: [4, 3, 5] }, { data: [1, 6, 3] }, { data: [2, 5, 6] }]}*/}
          {/*      width={500}*/}
          {/*      height={300}*/}
          {/*    />*/}
          {/*  </Card>*/}
          {/*</CustomTabPanel>*/}
        </Box>
      </Paper>

    </div>
  );
}