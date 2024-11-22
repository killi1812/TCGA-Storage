(async () => {
  const ping = async () => {
    const resp = await fetch("/api/ping-minio");
    const data = await resp.json();
    console.log(data);
  };
  const btn = document.getElementById("ping");
  btn.addEventListener("click", ping);
  const img = document.getElementById("img");
  const form = document.getElementById("upload");
  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const filenameInput = form.querySelector('[name="filename"]');
    const filename = filenameInput ? filenameInput.value.trim() : "";
    const formData = new FormData(form);
    const response = await fetch(form.action || "/", {
      method: form.method || "POST",
      body: formData,
    });
    const name = filename.split("\\")[2];
    img.src = "/api/img/" + name;
  });
})();
