import classes from './HeaderSimple.module.css';
import {Container} from "@mantine/core";


export default function TimeSinkHeader() {

    return (
        <header className={classes.header}>
            <Container>
                <h1>Time Sink</h1>
            </Container>
        </header>
    )
}