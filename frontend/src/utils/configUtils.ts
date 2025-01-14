
export const componentToConfigMap = {
    "second(s)": "s",
    "minute(s)": "m",
    "hour(s)": "h",
}

export const configToComponentMap = {
    "s": "Second(s)",
    "m": "Minutes(s)",
    "h": "Hour(s)"
}

export function getDurationFromConfig(checkInterval: string): any {

    return checkInterval.split(" ")[0]
}

export function getComponentQualifier(checkInterval: string): string {
    const parts: string[] = checkInterval.split(" ")

    //@ts-ignore
    return configToComponentMap[parts[1]];
}





