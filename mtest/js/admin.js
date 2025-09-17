let bookings = JSON.parse(localStorage.getItem("bookings")) || [];
let deletedBooking = null;
let dateSortAsc = true;
let countySortAsc = true;

// ✅ Render bookings
function renderBookings() {
  const tbody = document.querySelector("#bookingsTable tbody");
  tbody.innerHTML = "";
  bookings.forEach((b, i) => {
    const row = document.createElement("tr");
    row.innerHTML = `
      <td>${b.date || ""}</td>
      <td>${b.time || ""}</td>
      <td>${b.name || ""}</td>
      <td>${b.email || ""}</td>
      <td>${b.phone || ""}</td>
      <td>${b.service || ""}</td>
      <td>${b.region || ""}</td>
      <td><button onclick="deleteBooking(${i})"><i class="fas fa-trash"></i> Delete</button></td>
    `;
    tbody.appendChild(row);
  });
}

// ✅ Delete booking with undo support
function deleteBooking(index) {
  deletedBooking = { booking: bookings[index], index };
  bookings.splice(index, 1);
  localStorage.setItem("bookings", JSON.stringify(bookings));
  renderBookings();
}

function undoDelete() {
  if (deletedBooking) {
    bookings.splice(deletedBooking.index, 0, deletedBooking.booking);
    localStorage.setItem("bookings", JSON.stringify(bookings));
    renderBookings();
    deletedBooking = null;
  } else {
    alert("No booking to undo.");
  }
}

// ✅ Refresh
function refreshBookings() {
  bookings = JSON.parse(localStorage.getItem("bookings")) || [];
  renderBookings();
}

// ✅ Sort by Date with toggle + arrow
function sortByDate() {
  bookings.sort((a, b) => {
    const da = new Date(a.date);
    const db = new Date(b.date);
    return dateSortAsc ? da - db : db - da;
  });

  document.getElementById("dateArrow").textContent = dateSortAsc ? "🔼" : "🔽";
  dateSortAsc = !dateSortAsc;

  localStorage.setItem("bookings", JSON.stringify(bookings));
  renderBookings();
}

// ✅ Sort by County with toggle + arrow
function sortByCounty() {
  bookings.sort((a, b) => {
    const ra = (a.region || "").toLowerCase();
    const rb = (b.region || "").toLowerCase();
    return countySortAsc ? ra.localeCompare(rb) : rb.localeCompare(ra);
  });

  document.getElementById("countyArrow").textContent = countySortAsc ? "🔼" : "🔽";
  countySortAsc = !countySortAsc;

  localStorage.setItem("bookings", JSON.stringify(bookings));
  renderBookings();
}

// ✅ Print
function printBookings() {
  const printContent = document.getElementById("bookingsTable").outerHTML;
  const win = window.open("", "", "height=700,width=900");
  win.document.write("<html><head><title>Print Bookings</title></head><body>");
  win.document.write(printContent);
  win.document.write("</body></html>");
  win.document.close();
  win.print();
}

// ✅ Download CSV
function downloadCSV() {
  let csv = "Date,Time,Name,Email,Phone,Service,Region\n";
  bookings.forEach(b => {
    csv += `${b.date},${b.time || ""},${b.name},${b.email},${b.phone},${b.service},${b.region}\n`;
  });
  const blob = new Blob([csv], { type: "text/csv" });
  const link = document.createElement("a");
  link.href = URL.createObjectURL(blob);
  link.download = "bookings.csv";
  link.click();
}

// ✅ Download Excel
function downloadExcel() {
  let table = document.getElementById("bookingsTable").outerHTML;
  const blob = new Blob([table], { type: "application/vnd.ms-excel" });
  const link = document.createElement("a");
  link.href = URL.createObjectURL(blob);
  link.download = "bookings.xls";
  link.click();
}

// ✅ Download PDF (simple HTML → PDF)
function downloadPDF() {
  const printContent = document.getElementById("bookingsTable").outerHTML;
  const blob = new Blob([printContent], { type: "application/pdf" });
  const link = document.createElement("a");
  link.href = URL.createObjectURL(blob);
  link.download = "bookings.pdf";
  link.click();
}

// ✅ Share (basic share API)
function shareBookings() {
  if (navigator.share) {
    navigator.share({
      title: "Bookings Data",
      text: "Here are the latest bookings.",
      url: window.location.href
    }).catch(err => console.error("Error sharing:", err));
  } else {
    alert("Sharing not supported in this browser.");
  }
}

// ✅ Logout
function adminLogout() {
  localStorage.removeItem("isAdmin");
  alert("You have been logged out.");
  window.location.href = "index.html";
}

// ✅ Init
renderBookings();