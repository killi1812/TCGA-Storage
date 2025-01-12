(async () => {
  const ping = async () => {
    const resp = await fetch("/api/ping-minio");
    const data = await resp.json();
    alert(data);
  };

  const scrape = async () => {
    const resp = await fetch("/api/scrape");
    const data = await resp.json();
    alert(data);
  };

  const pingBtn = document.getElementById("ping");
  pingBtn.addEventListener("click", ping);

  const scrapeBtn = document.getElementById("scrape");
  scrapeBtn.addEventListener("click", scrape);

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

  const dataForm = document.getElementById("data");

  dataForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    const filenameInput = dataForm.querySelector('[name="filename"]');
    const filename = filenameInput ? filenameInput.value.trim() : "";
    const dataFormData = new FormData(dataForm);
    const response = await fetch(dataForm.action || "/", {
      method: dataForm.method || "POST",
      body: dataFormData,
    });
  });
})();
