import React from 'react';
import {Bar, BarChart, CartesianGrid, Legend, ResponsiveContainer, Tooltip, XAxis, YAxis} from 'recharts';


export default function UsageBar(props: any) {
  const data = props.data;
  const xLabel = props.xLabel;
  const yLabel = props.yLabel;

  const staticData = [
    {
      name: "Test",
      duration: 100
    },
    {
      name: "Test 2",
      duration: 200
    }
  ]

  return (
    <>
        <BarChart
          data={data}
          width={800}
          height={400}
        >
          <CartesianGrid/>
          <XAxis dataKey="Name"/>
          <YAxis dataKey="Duration"/>
          <Tooltip/>
          <Legend/>
          <Bar dataKey={"Duration"} fill={"#1976d2"}/>
        </BarChart>
    </>
  )


}