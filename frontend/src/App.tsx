import {useEffect, useState} from 'react';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";
import {GetDailyProcesses} from "../wailsjs/go/bindings/DataBinding";
import {dateToStandardString} from "./utils/timeUtils";

import "@mantine/core/styles.css";
import {MantineProvider} from "@mantine/core";
import {theme} from "./theme";
import {UsageBarChart} from "./components/UsageBarChart";


interface UsageInfo {
    name: string;
    seen: string;
    duration: string;
}

function App() {
    const [usageData, setUsageData] = useState<UsageInfo[]>();

    function getDailyUsage() {
        const formattedDateString = dateToStandardString(new Date());
        GetDailyProcesses(formattedDateString).then(data => {
            setUsageData(data);
        })
    }

    //note to future stephen; save yourself an hour of debugging. `[]` makes useEffect fire only once.
    useEffect(() => {
        getDailyUsage();
    }, [])

    return (
        <MantineProvider defaultColorScheme={"dark"} theme={theme}>
            <div id="App">
                <h1>Time Sink</h1>
            </div>
        </MantineProvider>
    )
}

export default App
