(async () => {
  const deleteAll = async () => {
    const resp = await fetch("/api/patients/delete");
    debugger;
    if (resp.ok) {
      alert("Succesfully deleted all patients");
    }
  };

  const scrape = async () => {
    const resp = await fetch("/api/scrape");
    if (resp.ok) {
      alert("Scraping job started");
    } else {
      const data = await resp.json();
      alert(data);
    }
  };

  const pingBtn = document.getElementById("delete");
  pingBtn.addEventListener("click", deleteAll);

  const scrapeBtn = document.getElementById("scrape");
  scrapeBtn.addEventListener("click", scrape);

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
    if (response.ok) {
      alert(await response.json());
    } else {
      alert(response.statusText);
    }
  });
})();
