(async () => {
  const resp = await fetch("/api");
  const data = await resp.json();
  console.log(data);
})();
