import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";
import {GetDailyProcesses} from "../wailsjs/go/bindings/DataBinding";
import {dateToStandardString} from "./utils/timeUtils";

function App() {
    const [resultText, setResultText] = useState("Please enter your name below 👇");
    const [name, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    const [usageData, setUsageData] = useState<any>();

    function greet() {
        Greet(name).then(updateResultText);
    }

    function getDailyUsage() {
        const formattedDateString = dateToStandardString(new Date());
        GetDailyProcesses(formattedDateString).then(data => {
            console.info(data);
        })

    }

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo"/>
            <div id="result" className="result">{resultText}</div>
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text"/>
                <button className="btn" onClick={greet}>Greet</button>
                <button className="btn" onClick={getDailyUsage}>TEST</button>
            </div>
            <div>
                {usageData}
            </div>
        </div>
    )
}

export default App
