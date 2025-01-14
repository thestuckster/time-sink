import {GetConfig} from "../../wailsjs/go/bindings/TimeSinkConfigBinding";
import {useEffect, useState} from "react";
import React from 'react';
import {bindings} from "../../wailsjs/go/models";
import ConfigDto = bindings.ConfigDto;
import {Button, Card, Form, InputNumber, InputNumberProps, Layout, Modal, Segmented, Table} from "antd";
import {componentToConfigMap, configToComponentMap, getComponentQualifier, getDurationFromConfig} from "../utils/configUtils";
import type { TableColumnsType, TableProps } from 'antd';
import {GetCurrentlyRunningProcesses} from "../../wailsjs/go/bindings/ProcessBinding";
import ProcessDto = bindings.ProcessDto;



function toConfigScheduleFormat(duration: number, qualifier: string) {
    // @ts-ignore
    return `${duration} ${componentToConfigMap[qualifier]}`;
}


interface DataType {
    key: React.Key;
    name: string
}

function buildTableData(applications : string[]) : DataType[]{
    let dataTypes : DataType[] = [];
    for (let i = 0; i < applications.length; i++) {
        const application = applications[i];
        const key = `${i}`;
        dataTypes.push({key: key, name: application});
    }

    return dataTypes;
}

const columns : TableColumnsType<DataType> = [
    {
        title: 'Application Name',
        dataIndex: 'name',
    },
    {
        title: 'Action',
        dataIndex: '',
        key: 'x',
        //TODO: delete button functionality
        render: () => <Button onClick={() => console.log("TODO remove")}>Remove</Button>
    }
]

export default function Config() {
    const [config, setConfig] = useState<ConfigDto>();
    const [duration, setDuration] = useState<number | string>();
    const [qualifier, setQualifier] = useState<string>();
    const [tableData, setTableData] = useState<DataType[]>([]);
    const [open, setOpen] = useState<boolean>(false);
    const [runningApplications, setRunningApplications] = useState<ProcessDto[]>([]);

    const timeOptions: string[] =['Second(s)', 'Minute(s)', 'Hour(s)']

    useEffect(() => {
        console.log("frontend get config call")
        GetConfig().then((res) => {
            console.log("Get Config Result ", res);
            setConfig(res);

            console.log("Config check interval ", res.check_interval);
            setDuration(getDurationFromConfig(res.check_interval));
            console.log("Duration: ", duration);

            const componentQualifier = getComponentQualifier(res.check_interval);
            console.log("Qualifier: ", componentQualifier);
            setQualifier(componentQualifier);

            setTableData(buildTableData(res.applications));
        });

        GetCurrentlyRunningProcesses().then(res => {
           setRunningApplications(res.processes);
        });

    }, []);

    const showModal = () => {
        setOpen(!open);
    }

    const onModalOk = () => {
        setOpen(!open);
    }

    const onModalCancel = () => {
        setOpen(false);
    }

    const onTimeSelect = (selected: string): void => {
        setQualifier(selected.toLowerCase());
    }

    const onDurationChange: InputNumberProps['onChange'] = (value) => {
        const v = value as string;
        setDuration(parseInt(v));
        console.log(duration);
    }

    const saveConfig = () => {
    }

    return(
        <div>
            <Card title={"Edit Configuration"} style={{backgroundColor: "white"}}>

                <Layout style={{backgroundColor: "white", textAlign: "center"}}>
                    <div>
                    <h2>Update Interval</h2>
                     <p>Update Every: </p>
                        <InputNumber defaultValue={duration} onChange={onDurationChange}/>
                        <Segmented options={timeOptions} defaultValue={qualifier} onChange={onTimeSelect}/>
                    </div>
                </Layout>
                <Layout style={{backgroundColor: "white", textAlign: "center"}}>
                    {/*TODO: center table on page*/}
                    <div style={{textAlign: "center"}}>
                        <h2>Applications</h2>
                        <Button type={"primary"} style={{marginBottom: "10px"}} onClick={showModal}>Save
                            Add New Application(s)
                        </Button>
                        <Table
                            style={{width: '50%', textAlign: "center"}}
                            size={"middle"}
                            bordered={true}
                            rowHoverable={true}
                            dataSource={tableData}
                            columns={columns}
                            virtual/>

                    </div>
                </Layout>
                <div>
                    <Modal
                        title={"Add New Application(s) to Time Sink Tracking"}
                        open={open}
                        onOk={onModalOk}
                        onCancel={onModalCancel}
                    >

                        <p>Select the currently running process(es) you'd like to track with Time Sink</p>
                        <br/>
                        {/*TODO: a whole new table for process selection*/}
                        <Table
                            style={{width: '50%', textAlign: "center"}}
                            size={"middle"}
                            bordered={true}
                            rowHoverable={true}
                            dataSource={tableData}
                            columns={columns}
                            virtual/>
                    </Modal>
                </div>

            </Card>
        </div>
    );
}