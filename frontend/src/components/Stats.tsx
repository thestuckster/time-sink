import React from 'react';
import {Content, Header} from "antd/es/layout/layout";
import {GetProp, Layout, Menu, MenuProps} from "antd";
import Sider from "antd/es/layout/Sider";

type MenuItem = GetProp<MenuProps, 'items'>[number];

const items: MenuItem[] = [
    {
        key: '1',
        label: "Daily Usage"
    },
    {
        key: '2',
        label: "Last 30 Days"
    },
    {
        key: '3',
        label: "Yearly Breakdown"
    }
]

export default function Stats() {
    const [selected, setSelected] = React.useState<string[]>(['1']);
    const [currentContent, setCurrentContent] = React.useState<string>('');

    // @ts-ignore
    const onMenuSelected = ({item, key, keyPath, selectedKeys, domEvent}) => {
        setSelected(selectedKeys);

        // @ts-ignore
        setCurrentContent(items[key-1]!['label']);
    }

    const renderContent = () => {
        if (currentContent === "Daily Usage") {
            return (<div><h1>Daily Usage TODO</h1></div>)
        } else if (currentContent === "Last 30 Days") {
            return (<div><h1>Last 30 Days</h1></div>)
        } else if (currentContent === "Yearly Breakdown") {
            return (<div><h1>Yearly Breakdown</h1></div>)
        }
    }

    return (
      <div>
          <Layout>
              <Content>
                  <Layout style={{padding: "20px 0", borderRadius: "10px"}}>
                      <Sider width={200}>
                          <Menu
                              mode={"vertical"}
                              defaultSelectedKeys={selected}
                              items={items}
                          />
                      </Sider>
                      <Content> {renderContent()}</Content>
                  </Layout>
              </Content>
          </Layout>
      </div>
    );
}