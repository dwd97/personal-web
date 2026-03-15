# About

I build apps, write about math and programming.

## GitHub Activity
![GitHub activity graph](https://ghchart.rshah.org/dwd97)

## Habit streak
<script src="https://d3js.org/d3.v7.min.js"></script>
<script src="https://unpkg.com/cal-heatmap/dist/cal-heatmap.min.js"></script>
<link rel="stylesheet" href="https://unpkg.com/cal-heatmap/dist/cal-heatmap.css">

<div id="heatmap-scroll-wrapper" style="overflow-x: auto; overflow-y: hidden; padding-bottom: 15px; margin: 28px 0; scrollbar-width: thin;">
  <div id="todo-heatmap" style="min-width: 800px;"></div>
</div>

<script>
  let cal = new CalHeatmap();

  function renderHeatmap() {
    const isDark = document.documentElement.getAttribute('data-theme') === 'dark';
    const emptyColor = isDark ? '#161b22' : '#ebedf0';

    // Strictly align the start date to the 1st of the month, 12 months prior
    const today = new Date();
    const startDate = new Date(today.getFullYear(), today.getMonth() - 12, 1);

    cal.paint({
      itemSelector: '#todo-heatmap',
      domain: { 
        type: 'month', 
        gutter: 0, 
        label: { text: 'MMM', textAlign: 'start', position: 'top' } 
      },
      subDomain: { 
        type: 'ghDay', 
        radius: 2, 
        width: 11, 
        height: 11, 
        gutter: 4 
      },
      range: 13, // 12 historical months + the current active month
      date: { start: startDate },
      data: { 
        source: '/habits.json', 
        type: 'json', 
        x: 'date', 
        y: 'value' 
      },
      scale: {
        color: {
          type: 'threshold',
          range: [emptyColor, '#39d353'],
          domain: [1]
        }
      }
    }).then(() => {
      // Execute scroll alignment post-render to display the rightmost edge
      const scrollWrapper = document.getElementById('heatmap-scroll-wrapper');
      if (scrollWrapper) {
        scrollWrapper.scrollLeft = scrollWrapper.scrollWidth;
      }
    });
  }

  renderHeatmap();

  const observer = new MutationObserver((mutations) => {
    mutations.forEach((mutation) => {
      if (mutation.attributeName === 'data-theme') {
        cal.destroy().then(() => {
          cal = new CalHeatmap();
          renderHeatmap();
        });
      }
    });
  });

  observer.observe(document.documentElement, { attributeFilter: ['data-theme'] });
</script>