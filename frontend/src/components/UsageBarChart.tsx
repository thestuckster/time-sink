import {BarChart} from "@mantine/charts";

export interface UsageChartSeries {
    name: string;
    color: string;
}

export interface UsageBarChartProps {
    height: number;
    data: any;
    dataKey: string;
    series: UsageChartSeries[];
    tickLine: "y" | "x";
    withLegend: boolean;
}

export function UsageBarChart({
    height, data, dataKey, series, tickLine, withLegend,
                       }: UsageBarChartProps) {

    return (
        <div>TEST</div>
    )
}