# About

I build apps, write about math and programming.

<h2>GitHub Activity</h2>

<div id="github-graph"></div>

<script>
async function loadGitHubGraph() {
  const res = await fetch("https://github-contributions-api.jogruber.de/v4/dwd97?y=last");
  const data = await res.json();

  const container = document.getElementById("github-graph");

  const grid = document.createElement("div");
  grid.style.display = "grid";
  grid.style.gridTemplateColumns = "repeat(53, 10px)";
  grid.style.gap = "3px";

  data.contributions.forEach(day => {
    const cell = document.createElement("div");
    cell.style.width = "10px";
    cell.style.height = "10px";

    const colors = [
      "#ebedf0",
      "#9be9a8",
      "#40c463",
      "#30a14e",
      "#216e39"
    ];

    cell.style.background = colors[day.level];
    grid.appendChild(cell);
  });

  container.appendChild(grid);
}

loadGitHubGraph();
</script>