(async () => {
  const ping = async () => {
    const resp = await fetch("/api/ping-minio");
    const data = await resp.json();
    console.log(data);
  };
  const btn = document.getElementById("ping");
  btn.addEventListener("click", ping);
})();
