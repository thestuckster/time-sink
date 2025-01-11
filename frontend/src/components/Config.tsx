import {GetConfig} from "../../wailsjs/go/bindings/TimeSinkConfigBinding";
import {useEffect, useState} from "react";
import React from 'react';
import {bindings} from "../../wailsjs/go/models";
import ConfigDto = bindings.ConfigDto;
import {Greet} from "../../wailsjs/go/main/App";
import {Card, InputNumber, InputNumberProps, Layout, Segmented} from "antd";

function toConfigScheduleFormat(duration: number, qualifier: string) {
    const mappings = {
        "second(s)": "s",
        "minute(s)": "m",
        "hour(s)": "h",
    }

    // @ts-ignore
    return `${duration} ${mappings[qualifier]}`;
}

export default function Config() {
    const [config, setConfig] = useState<ConfigDto>();
    const [duration, setDuration] = useState<number>();
    const [qualifier, setQualifier] = useState<string>();

    const timeOptions: string[] =['Second(s)', 'Minute(s)', 'Hour(s)']

    useEffect(() => {
        console.log("frontend get config call")
        GetConfig().then((res) => {
            console.log("Get Config Result ", res);
            setConfig(res);
        });
    }, [config]);


    const onTimeSelect = (selected: string): void => {
        setQualifier(selected.toLowerCase());
    }

    const onDurationChange: InputNumberProps['onChange'] = (value) => {
        console.log(value);
    }

    const saveConfig = () => {

    }

    return(
        <div>
            <Card title={"Edit Configuration"} style={{backgroundColor: "white", textAlign: "center"}}>
                <Layout style={{backgroundColor: "white", textAlign: "center"}}>
                    <p>Update Every:
                        <InputNumber defaultValue={duration} onChange={onDurationChange}/>
                        <Segmented options={timeOptions} onChange={onTimeSelect}/>
                    </p>
                </Layout>
            </Card>
        </div>
    );
}