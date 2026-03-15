# About

I build apps, write about math and programming.

## GitHub Activity
![GitHub activity graph](https://ghchart.rshah.org/dwd97)

## Habit streak
<script src="https://d3js.org/d3.v7.min.js"></script>
<script src="https://unpkg.com/cal-heatmap/dist/cal-heatmap.min.js"></script>
<link rel="stylesheet" href="https://unpkg.com/cal-heatmap/dist/cal-heatmap.css">

<div id="todo-heatmap"></div>

<script>
  const cal = new CalHeatmap();
  
  // Calculate start date (11 months prior to current month)
  const startDate = new Date();
  startDate.setMonth(startDate.getMonth() - 11);

  cal.paint({
    itemSelector: '#todo-heatmap',
    domain: { type: 'month' },
    subDomain: { type: 'day', radius: 2, width: 11, height: 11, gutter: 4 },
    date: { start: startDate },
    data: {
      source: '/habits.json',
      type: 'json',
      x: 'date',
      y: 'value'
    },
    // Add the scale configuration here
    scale: {
      color: {
        type: 'threshold',
        range: ['#ebedf0', '#39d353'], // 0 renders as #ebedf0 (empty), 1 renders as #39d353 (green)
        domain: [1] 
      }
    },
    theme: 'light'
  });
</script>