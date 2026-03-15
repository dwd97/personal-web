# About

I build apps, write about math and programming.

## GitHub Activity
![GitHub activity graph](https://ghchart.rshah.org/dwd97)

## Habit streak

<style>
  .heatmap-wrapper {
    display: flex;
    margin: 28px 0;
    font-family: Inter, system-ui, sans-serif;
    color: var(--muted);
    font-size: 11px;
    user-select: none;
  }

  /* Y-Axis (Days) */
  .heatmap-y {
    display: grid;
    grid-template-rows: repeat(7, 10px);
    gap: 3px;
    margin-top: 20px; /* Aligns with grid, leaves space for X-axis */
    padding-right: 8px;
  }
  .heatmap-y span {
    line-height: 10px;
  }
  .heatmap-y span:nth-child(1) { grid-row: 2; } /* Monday */
  .heatmap-y span:nth-child(2) { grid-row: 4; } /* Wednesday */
  .heatmap-y span:nth-child(3) { grid-row: 6; } /* Friday */

  /* Scrollable Area */
  .heatmap-scroll {
    overflow-x: auto;
    scrollbar-width: none; /* Firefox */
  }
  .heatmap-scroll::-webkit-scrollbar {
    display: none; /* Chrome/Safari */
  }

  /* X-Axis (Months) */
  .heatmap-x {
    position: relative;
    height: 20px;
  }
  .heatmap-x span {
    position: absolute;
    bottom: 4px;
  }

  /* Grid */
  .heatmap-grid {
    display: grid;
    grid-template-rows: repeat(7, 10px);
    grid-auto-flow: column;
    gap: 3px;
  }

  /* Cells */
  .heatmap-cell {
    width: 10px;
    height: 10px;
    border-radius: 2px;
    background-color: #ebedf0; /* Light Mode Empty */
    outline: 1px solid rgba(27, 31, 35, 0.06);
    outline-offset: -1px;
  }
  [data-theme="dark"] .heatmap-cell {
    background-color: #161b22; /* Dark Mode Empty */
    outline: 1px solid rgba(255, 255, 255, 0.05);
  }
  .heatmap-cell.active {
    background-color: #39d353 !important;
    outline: none;
  }
</style>

<div class="heatmap-wrapper">
  <div class="heatmap-y">
    <span>Mon</span>
    <span>Wed</span>
    <span>Fri</span>
  </div>
  <div class="heatmap-scroll" id="heatmap-scroll">
    <div class="heatmap-x" id="heatmap-months"></div>
    <div class="heatmap-grid" id="heatmap-squares"></div>
  </div>
</div>

<script>
  async function renderHeatmap() {
    const grid = document.getElementById('heatmap-squares');
    const months = document.getElementById('heatmap-months');
    const scroll = document.getElementById('heatmap-scroll');

    // 1. Data Ingestion
    let habitMap = new Map();
    try {
      const res = await fetch('/habits.json');
      const data = await res.json();
      data.forEach(d => habitMap.set(d.date, d.value));
    } catch (err) {
      console.warn("Failed to load /habits.json");
    }

    // 2. Date Domain Calculation (52 weeks prior to today)
    const endDate = new Date();
    endDate.setHours(0, 0, 0, 0);

    const startDate = new Date(endDate);
    startDate.setDate(endDate.getDate() - (52 * 7) - endDate.getDay());

    let current = new Date(startDate);
    let squaresHtml = '';
    let monthsHtml = '';
    let colIndex = 0;

    // 3. Grid Generation
    while (current <= endDate) {
      // Month label placement (triggers on the 1st of the month)
      if (current.getDate() === 1 && colIndex > 0) {
        const monthName = current.toLocaleString('en-US', { month: 'short' });
        // 13px offset per column (10px width + 3px gap)
        monthsHtml += `<span style="left: ${colIndex * 13}px">${monthName}</span>`;
      }

      // Date formatting to YYYY-MM-DD
      const year = current.getFullYear();
      const month = String(current.getMonth() + 1).padStart(2, '0');
      const day = String(current.getDate()).padStart(2, '0');
      const dateStr = `${year}-${month}-${day}`;

      // Value mapping
      const isActive = habitMap.get(dateStr) >= 1;
      const cssClass = isActive ? 'active' : '';

      squaresHtml += `<div class="heatmap-cell ${cssClass}" title="${dateStr}"></div>`;

      // Advance iterator
      current.setDate(current.getDate() + 1);
      if (current.getDay() === 0) {
        colIndex++;
      }
    }

    // 4. DOM Injection & Scroll Alignment
    grid.innerHTML = squaresHtml;
    months.innerHTML = monthsHtml;
    scroll.scrollLeft = scroll.scrollWidth;
  }

  renderHeatmap();
</script>