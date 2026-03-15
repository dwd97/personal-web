# About

I build apps, write about math and programming.

## GitHub Activity
![GitHub activity graph](https://ghchart.rshah.org/dwd97)

## Habit streak
<script src="https://d3js.org/d3.v7.min.js"></script>
<script src="https://unpkg.com/cal-heatmap/dist/cal-heatmap.min.js"></script>
<link rel="stylesheet" href="https://unpkg.com/cal-heatmap/dist/cal-heatmap.css">

<div class="code-scroll" style="margin: 28px 0;">
  <div id="todo-heatmap" style="min-width: 780px;"></div>
</div>

<script>
  let cal = new CalHeatmap();
  const today = new Date();
  const startDate = new Date(today.getFullYear(), today.getMonth() - 11, 1);

function renderHeatmap() {
    const isDark = document.documentElement.getAttribute('data-theme') === 'dark';
    const emptyColor = isDark ? '#161b22' : '#ebedf0';

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
      range: 12,
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
      // Execute scroll alignment post-render
      const heatmapContainer = document.getElementById('todo-heatmap');
      const scrollWrapper = heatmapContainer.parentElement;
      
      if (scrollWrapper) {
        // Force the scrollbar to the maximum right position
        scrollWrapper.scrollLeft = scrollWrapper.scrollWidth;
      }
    });
  }

  // Initial execution
  renderHeatmap();

  // Mutation observer for dynamic theme switching
  const observer = new MutationObserver((mutations) => {
    mutations.forEach((mutation) => {
      if (mutation.attributeName === 'data-theme') {
        // Destroy existing instance, then create a new one
        cal.destroy().then(() => {
          cal = new CalHeatmap();
          renderHeatmap();
        });
      }
    });
  });

  observer.observe(document.documentElement, { attributeFilter: ['data-theme'] });
</script>