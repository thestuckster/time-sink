import {GetConfig} from "../../wailsjs/go/bindings/TimeSinkConfigBinding";
import {useEffect, useState} from "react";
import React from 'react';
import {bindings} from "../../wailsjs/go/models";
import ConfigDto = bindings.ConfigDto;

export default function Config() {
    const [config, setConfig] = useState<ConfigDto>();

    useEffect(() => {
        GetConfig().then((config) => {
            setConfig(config)
        })
    }, [config])

    return(
        <div>
            <p>Check Interval: {config?.check_interval}</p>
        </div>
    )

}