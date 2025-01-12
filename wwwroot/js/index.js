(() => {
  let allPatients = [];
  let filteredPatients = [];
  let currentIndex = 0;
  const patientsPerPage = 20; // Number of patients to load at a time
  let isLoading = false;
  let isFirstLoad = true; // Track if it's the first page load

  // Fetch patients from API and populate the table
  async function fetchPatients() {
    const loadingSpinner = document.getElementById("loadingSpinner");
    const tableBody = document.getElementById("patientsTableBody");

    // Show the loading spinner
    loadingSpinner.style.display = "flex";

    try {
      const response = await fetch("/api/patients"); // Fetch from endpoint
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      allPatients = await response.json(); // Parse JSON response
      filteredPatients = [...allPatients]; // Initialize filteredPatients
      displayPatients();
    } catch (error) {
      console.error("Error fetching patients:", error);
      tableBody.innerHTML = `<tr><td colspan="4" class="text-center text-danger">Failed to load data</td></tr>`;
    } finally {
      // Hide the loading spinner
      loadingSpinner.style.display = "none";
    }
  }

  // Function to display patients in the table (pagination)
  function displayPatients() {
    const tableBody = document.getElementById("patientsTableBody");
    const startIndex = currentIndex;
    const endIndex = Math.min(
      currentIndex + patientsPerPage,
      filteredPatients.length,
    );
    const patientsToDisplay = filteredPatients.slice(startIndex, endIndex);

    // Use a single fragment to append all rows at once for better performance
    const fragment = document.createDocumentFragment();
    patientsToDisplay.forEach((patient) => {
      const row = document.createElement("tr");
      row.onclick = () => navigateToPatientDetails(patient.bcr_patient_barcode); // Add click event
      row.innerHTML = `
                    <td>${patient.bcr_patient_barcode}</td>
                    <td>${patient.dss ? "Yes" : "No"}</td>
                    <td>${patient.os ? "Yes" : "No"}</td>
                    <td>${patient.clinical_stage || "N/A"}</td>
                `;
      fragment.appendChild(row);
    });
    tableBody.appendChild(fragment); // Append all rows at once

    // Update current index for next batch of data
    currentIndex = endIndex;

    // Check if we have loaded all patients and disable further loading
    if (currentIndex >= filteredPatients.length) {
      window.removeEventListener("scroll", handleScroll);
    }

    // Reset scroll position to the top only on the first load
    if (isFirstLoad) {
      const tableContainer = document.getElementById("tableContainer");
      tableContainer.scrollTop = 0;
      isFirstLoad = false; // Ensure this only happens once
    }
  }

  // Handle search form submit
  function handleSearch(event) {
    event.preventDefault();
    const searchValue = document
      .getElementById("patientCodeSearch")
      .value.toLowerCase();
    filteredPatients = allPatients.filter((patient) => {
      return patient.bcr_patient_barcode.toLowerCase().includes(searchValue);
    });
    currentIndex = 0; // Reset to start from the beginning
    document.getElementById("patientsTableBody").innerHTML = ""; // Clear the table
    displayPatients(); // Display filtered patients
  }

  // Lazy loading: load more patients when scrolling to the bottom of the table container
  function handleScroll() {
    const tableContainer = document.getElementById("tableContainer");
    const bottom =
      tableContainer.scrollHeight ===
      tableContainer.scrollTop + tableContainer.clientHeight;
    if (bottom && !isLoading) {
      isLoading = true;
      displayPatients();
      isLoading = false;
    }
  }

  document
    .getElementById("searchForm")
    .addEventListener("submit", handleSearch);

  // Navigate to patient details page
  function navigateToPatientDetails(patientBarcode) {
    window.location.href = `patient-details/${patientBarcode}`; // Adjust this URL as needed
  }

  // Fetch patients when the page loads
  window.onload = () => {
    fetchPatients(); // Fetch all patients
    const tableContainer = document.getElementById("tableContainer");
    tableContainer.addEventListener("scroll", handleScroll); // Add event listener for lazy loading
  };
})();
