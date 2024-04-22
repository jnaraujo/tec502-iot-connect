import ApexChart from "react-apexcharts"
import { teal } from "tailwindcss/colors"

interface Props {
  data: number[]
  title: string
  className?: string
}

export function Chart(props: Props) {
  return (
    <section className={props.className}>
      <ApexChart
        type="area"
        width="100%"
        height="100%"
        options={{
          chart: {
            id: "webhook-events-amount-chart",
            toolbar: {
              show: false,
            },
            animations: {
              easing: "linear",
              dynamicAnimation: {
                speed: 1000,
              },
            },
            parentHeightOffset: 0,
            sparkline: {
              enabled: false,
            },
            zoom: {
              enabled: false,
            },
          },
          grid: {
            show: false,
            padding: {
              left: -9,
              right: -1,
              bottom: -8,
              top: -20,
            },
          },
          tooltip: {
            enabled: false,
          },
          colors: [teal[400]],
          stroke: {
            curve: "smooth",
            width: 2,
            lineCap: "butt",
          },
          fill: {
            gradient: {
              opacityFrom: 0.8,
              opacityTo: 0.4,
            },
          },
          dataLabels: {
            enabled: false,
          },
          xaxis: {
            labels: {
              show: false,
            },
            axisTicks: {
              show: false,
            },
            axisBorder: {
              show: false,
            },
            tooltip: {
              enabled: false,
            },
          },
          yaxis: {
            labels: {
              show: false,
            },
          },
        }}
        series={[
          {
            name: props.title,
            data: props.data,
          },
        ]}
      />
    </section>
  )
}
