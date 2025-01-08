import {useEffect, useState} from 'react';
import './App.css';

import "@mantine/core/styles.css";
import {AppShell, Burger, MantineProvider} from "@mantine/core";
import {theme} from "./theme";
import {useDisclosure} from "@mantine/hooks";


interface UsageInfo {
    name: string;
    seen: string;
    duration: number;
}

interface Processes {
    processes: string[];
}

function App() {

    const [opened, { toggle }] = useDisclosure();

    return (
        <MantineProvider defaultColorScheme={"dark"} theme={theme}>
            <div id="App">
                <h1>Time Sink</h1>
                <AppShell
                    header={{ height: 60 }}
                    navbar={{
                        width: 300,
                        breakpoint: 'sm',
                        collapsed: { mobile: !opened },
                    }}
                    padding="md"
                >
                    <AppShell.Header>
                        <Burger
                            opened={opened}
                            onClick={toggle}
                            hiddenFrom="sm"
                            size="sm"
                        />
                        <div>Logo</div>
                    </AppShell.Header>

                    <AppShell.Navbar p="md">Navbar</AppShell.Navbar>

                    <AppShell.Main>Main</AppShell.Main>
                </AppShell>
            </div>
        </MantineProvider>
    );
}

export default App
