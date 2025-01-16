import {Box, Card, Paper, Tab, Tabs, Typography} from "@mui/material";
import {useState} from "react";
import {BarChart} from "@mui/x-charts";
import {bindings} from "../../wailsjs/go/models";
import ProcessUsageData = bindings.ProcessUsageData;


//TODO: delete

function generateDummyDailyUsage() : ProcessUsageData[]{
  return [
    {
      name: "TestApplication",
      duration: 100,
      seen: 100
    },
    {
      name: "Another application",
      duration: 34,
      seen: 100
    }
  ]
}

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

function CustomTabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && <Box sx={{ p: 3 }}>{children}</Box>}
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

  const onTabChange = (event: React.SyntheticEvent, tab: number): void => {
    setTabValue(tab);
  }

  return (
    <>
     <Paper elevation={4}>
       <Box>
         <Tabs value={tabValue} onChange={onTabChange}>
           <Tab label={"Today's Usage"} {...tabProps(0)} />
           <Tab label={"Last 30 Days"}  {...tabProps(1)} />
           <Tab label={"Yearly Usage"}  {...tabProps(2)} />
         </Tabs>
         <CustomTabPanel index={0} value={tabValue}>
            <h1>Today</h1>
           <BarChart
             xAxis={[{ scaleType: 'band', data: ['group A', 'group B', 'group C'] }]}
             series={[{ data: [4, 3, 5] }, { data: [1, 6, 3] }, { data: [2, 5, 6] }]}
             width={500}
             height={300}
           />
         </CustomTabPanel>
         <CustomTabPanel index={1} value={tabValue}>
           <h1>Last 30</h1>
           <BarChart
             xAxis={[{ scaleType: 'band', data: ['group A', 'group B', 'group C'] }]}
             series={[{ data: [4, 3, 5] }, { data: [1, 6, 3] }, { data: [2, 5, 6] }]}
           />
         </CustomTabPanel>
         <CustomTabPanel index={2} value={tabValue}>
           <Card>
             <BarChart
               xAxis={[{ scaleType: 'band', data: ['group A', 'group B', 'group C'] }]}
               series={[{ data: [4, 3, 5] }, { data: [1, 6, 3] }, { data: [2, 5, 6] }]}
               width={500}
               height={300}
             />
           </Card>
         </CustomTabPanel>
       </Box>
     </Paper>

    </>
  );
}