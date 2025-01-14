import React from 'react';
import {GetProp, Layout, Menu, MenuProps} from 'antd';
import {PieChartOutlined, SettingOutlined} from "@ant-design/icons";
import {Content, Header} from "antd/es/layout/layout";
import Config from "./components/Config";
import Stats from "./components/Stats";

type MenuItem = GetProp<MenuProps, 'items'>[number];

interface UsageInfo {
    name: string;
    seen: string;
    duration: number;
}

interface Processes {
    processes: string[];
}

const items: MenuItem[] = [
    {
        key: '1',
        icon: <PieChartOutlined/>,
        label: "Stats"
    },
    {
        key: '2',
        icon: <SettingOutlined/>,
        label: "Config"
    }

]

function App() {

    const [selected, setSelected] = React.useState<string[]>(['1']);
    const [currentContent, setCurrentContent] = React.useState('');

    // @ts-ignore
    const onMenuSelected = ({item, key, keyPath, selectedKeys, domEvent}) => {
        setSelected(selectedKeys);

        // @ts-ignore
        setCurrentContent(items[key-1]!['label']);
    }

    const renderContent = () => {
        console.log("renderContent");
        if (currentContent == "Config") {
            return <Config/>
        } else {
            return <Stats/>;
        }
    }

    return (
        <>
            <Layout>
                <Header style={{display: 'flex', alignItems: 'center', background: "white"}}>
                    <div style={{backgroundColor:'white'}}>
                        Time Sink
                    </div>
                    <Menu
                        mode={"horizontal"}
                        defaultSelectedKeys={selected}
                        items={items}
                        style={{flex: 1, minWidth: 0}}
                        onSelect={onMenuSelected}
                    />
                </Header>
                <Content>
                    {renderContent()}
                </Content>
            </Layout>
        </>
    );
}

export default App
