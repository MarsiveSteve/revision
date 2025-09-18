// index.js - main client logic (counties, booking, admin accounts)

// wrap in DOMContentLoaded so elements exist
document.addEventListener('DOMContentLoaded', () => {

  // ---------- Accounts helpers ----------
  function loadAccounts() {
    return JSON.parse(localStorage.getItem('adminAccounts') || '{}');
  }
  function saveAccounts(obj) {
    localStorage.setItem('adminAccounts', JSON.stringify(obj));
  }

  // ensure default admin exists
  const existingAccounts = loadAccounts();
  if (!existingAccounts.admin) {
    existingAccounts.admin = "admin123";
    saveAccounts(existingAccounts);
  }

  // expose modal functions to global (onclick in HTML uses them)
  window.showAdminLogin = function() {
    const m = document.getElementById('adminLoginModal');
    m.classList.add('show'); m.style.display = 'flex';
    document.getElementById('loginError').textContent = '';
  };
  window.showCreateAccount = function() {
    const m = document.getElementById('createAccountModal');
    m.classList.add('show'); m.style.display = 'flex';
    document.getElementById('createMsg').textContent = '';
    // clear fields
    document.getElementById('newUser').value = '';
    document.getElementById('newPass').value = '';
    document.getElementById('confirmPass').value = '';
  };
  window.closeCreateAccount = function() {
    const m = document.getElementById('createAccountModal');
    m.classList.remove('show'); m.style.display = 'none';
  };
  window.showForgotPassword = function() {
    const m = document.getElementById('forgotPassModal');
    m.classList.add('show'); m.style.display = 'flex';
    document.getElementById('resetMsg').textContent = '';
    document.getElementById('resetUser').value = '';
    document.getElementById('resetPass').value = '';
    document.getElementById('confirmResetPass').value = '';
  };
  window.closeForgotPassword = function() {
    const m = document.getElementById('forgotPassModal');
    m.classList.remove('show'); m.style.display = 'none';
  };

  // ---------- Admin actions ----------
  window.adminLogin = function() {
    const user = document.getElementById('adminUser').value.trim();
    const pass = document.getElementById('adminPass').value.trim();
    const accounts = loadAccounts();
    if (accounts[user] && accounts[user] === pass) {
      localStorage.setItem('isAdmin','true');
      localStorage.setItem('activeAdmin', user);
      // redirect to admin.html (make sure admin.html exists)
      window.location.href = 'admin.html';
    } else {
      document.getElementById('loginError').textContent = 'Invalid username or password';
    }
  };

  window.createAccount = function() {
    const user = document.getElementById('newUser').value.trim();
    const pass = document.getElementById('newPass').value;
    const confirm = document.getElementById('confirmPass').value;

    if (!user || !pass || !confirm) {
      document.getElementById('createMsg').textContent = 'Please fill all fields';
      return;
    }
    if (pass !== confirm) {
      document.getElementById('createMsg').textContent = 'Passwords do not match';
      return;
    }
    const accounts = loadAccounts();
    if (accounts[user]) {
      document.getElementById('createMsg').textContent = 'Username already exists';
      return;
    }
    accounts[user] = pass;
    saveAccounts(accounts);
    document.getElementById('createMsg').textContent = 'Account created ✅';
    setTimeout(closeCreateAccount, 700);
  };

  window.resetPassword = function() {
    const user = document.getElementById('resetUser').value.trim();
    const pass = document.getElementById('resetPass').value;
    const confirm = document.getElementById('confirmResetPass').value;
    if (!user || !pass || !confirm) {
      document.getElementById('resetMsg').textContent = 'Please fill all fields';
      return;
    }
    if (pass !== confirm) {
      document.getElementById('resetMsg').textContent = 'Passwords do not match';
      return;
    }
    const accounts = loadAccounts();
    if (!accounts[user]) {
      document.getElementById('resetMsg').textContent = 'Username not found';
      return;
    }
    accounts[user] = pass;
    saveAccounts(accounts);
    document.getElementById('resetMsg').textContent = 'Password reset ✅';
    setTimeout(closeForgotPassword, 700);
  };

  // ---------- Counties & UI data ----------
  const messages = {
    "Nairobi": { text: "Corporate packages available for offices, estates & rooftop gardens.", icon: "https://img.icons8.com/emoji/96/deciduous-tree.png" },
    "Mombasa": { text: "Special coastal lawn care services with drought-resistant options.", icon: "https://img.icons8.com/emoji/96/palm-tree.png" },
    "Kisumu": { text: "Community tree planting hub around Lake Victoria.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Nakuru": { text: "Eco-friendly estate and gated community projects.", icon: "https://img.icons8.com/emoji/96/herb.png" },
    "Kiambu": { text: "Premium lawn care for residential and golf courses.", icon: "https://img.icons8.com/emoji/96/evergreen-tree.png" },
    "Kericho": { text: "Kericho tea estate greening initiatives.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Kajiado": { text: "Savannah lawn and garden services.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Machakos": { text: "Horticultural and residential garden care.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Embu": { text: "Community tree planting programs.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Meru": { text: "Green spaces development for estates.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Tharaka-Nithi": { text: "Rural greening and farm landscaping.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Kitui": { text: "Desert landscaping & drought-resistant lawns.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Makueni": { text: "Residential and corporate garden services.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Taita-Taveta": { text: "Coastal-safari region greening projects.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Kwale": { text: "Beachside lawns & community greening.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Lamu": { text: "Island gardens and heritage site landscaping.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Kilifi": { text: "Drought-tolerant lawns & trees.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Tana River": { text: "Riverine and farm landscaping.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Garissa": { text: "Arid region eco-lawn projects.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Wajir": { text: "Sustainable desert greenery initiatives.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Mandera": { text: "Community trees & green spaces.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Marsabit": { text: "Highland & desert lawn projects.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Isiolo": { text: "Eco-lawn programs for settlements.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Mwingi": { text: "Community horticulture projects.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Nyeri": { text: "Highland lawns & tree planting.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Kirinyaga": { text: "Residential and farm greening.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Murang'a": { text: "Corporate estates & green spaces.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Turkana": { text: "Arid-land greening programs.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "West Pokot": { text: "Community greening initiatives.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Samburu": { text: "Highland & riverine eco-lawn projects.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Trans Nzoia": { text: "Farm & estate landscaping services.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Uasin Gishu": { text: "Golf courses & residential lawns.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Elgeyo-Marakwet": { text: "Highland eco-gardening programs.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Nandi": { text: "Community greening and lawn care.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Baringo": { text: "River valleys and community lawns.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Laikipia": { text: "Wildlife-friendly landscaping projects.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Narok": { text: "Savannah estate & park greening.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Bomet": { text: "Community greening and lawns.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
    "Tanzania": { text: "Expanding to Dar & Arusha with corporate services.", icon: "https://img.icons8.com/emoji/96/flag-tanzania.png" },
    "Uganda": { text: "Launching Kampala green schools program.", icon: "https://img.icons8.com/emoji/96/flag-uganda.png" },
    "Rwanda": { text: "Eco-friendly Kigali greening initiatives.", icon: "https://img.icons8.com/emoji/96/flag-rwanda.png" },
    "Burundi": { text: "Community agroforestry & lawns in Bujumbura.", icon: "https://img.icons8.com/emoji/96/flag-burundi.png" }
  };

  const regions = Object.keys(messages);
  const counties = regions.filter(r => !["Tanzania","Uganda","Rwanda","Burundi"].includes(r));
  const eaCountries = ["Tanzania","Uganda","Rwanda","Burundi"];

  // pagination
  let currentPage = 1;
  const perPage = 8;

  function displayCounties(page = 1) {
    const grid = document.getElementById('countyGrid');
    grid.innerHTML = '';
    const start = (page - 1) * perPage;
    const end = start + perPage;
    counties.slice(start, end).forEach(c => {
      const card = document.createElement('div');
      card.className = 'card';
      card.textContent = c;
      card.onclick = () => showModal(c);
      grid.appendChild(card);
    });
    document.getElementById('pageInfo').textContent = `Page ${page} of ${Math.ceil(counties.length / perPage)}`;
  }

  document.getElementById('prevPage').onclick = () => {
    if (currentPage > 1) { currentPage--; displayCounties(currentPage); }
  };
  document.getElementById('nextPage').onclick = () => {
    if (currentPage < Math.ceil(counties.length / perPage)) { currentPage++; displayCounties(currentPage); }
  };

  // EA grid
  const eaGrid = document.getElementById('eaGrid');
  eaCountries.forEach(c => {
    const card = document.createElement('div');
    card.className = 'card';
    card.textContent = c;
    card.onclick = () => showModal(c);
    eaGrid.appendChild(card);
  });

  // populate region dropdown
  const regionSelect = document.getElementById('regionSelect');
  regions.forEach(r => {
    const opt = document.createElement('option');
    opt.value = r; opt.textContent = r;
    regionSelect.appendChild(opt);
  });

  // modal show
  window.showModal = function(region) {
    const info = messages[region];
    document.getElementById('countyName').textContent = region;
    document.getElementById('countyInfo').textContent = info.text;
    document.getElementById('regionIcon').src = info.icon;
    const m = document.getElementById('countyModal');
    m.classList.add('show'); m.style.display = 'flex';
  };
  window.closeModal = function() {
    const m = document.getElementById('countyModal');
    m.classList.remove('show'); m.style.display = 'none';
  };

  // search
  window.filterRegions = function() {
    const input = document.getElementById('searchBox').value.toLowerCase();
    const list = document.getElementById('regionList');
    list.innerHTML = '';
    if (!input) { list.style.display = 'none'; return; }
    regions.forEach(r => {
      if (r.toLowerCase().includes(input)) {
        const div = document.createElement('div');
        div.className = 'region-item'; div.textContent = r;
        div.onclick = () => {
          document.getElementById('searchBox').value = r;
          document.getElementById('regionSelect').value = r;
          showModal(r);
          list.style.display = 'none';
        };
        list.appendChild(div);
      }
    });
    list.style.display = list.children.length ? 'block' : 'none';
  };

  // booking form
  const bookingForm = document.getElementById('bookingForm');
  bookingForm.addEventListener('submit', e => {
    e.preventDefault();
    const bookings = JSON.parse(localStorage.getItem('bookings') || '[]');
    bookings.push({
      name: document.getElementById('fullName').value,
      email: document.getElementById('email').value,
      phone: document.getElementById('phone').value,
      service: document.getElementById('serviceType').value,
      region: document.getElementById('regionSelect').value,
      date: document.getElementById('preferredDate').value,
      time: new Date().toLocaleString()
    });
    localStorage.setItem('bookings', JSON.stringify(bookings));
    alert('✅ Booking submitted successfully!');
    bookingForm.reset();
    // if admin is logged in and viewing prondle update there
  });

  // initial render
  displayCounties(currentPage);

}); // end DOMContentLoaded

// Animate the tree count
let count = 0;
const target = 15847;
const treeCounter = document.getElementById("tree-count");

const interval = setInterval(() => {
  count += Math.ceil(target / 100); // speed factor
  if (count >= target) {
    count = target;
    clearInterval(interval);
  }
  treeCounter.textContent = count.toLocaleString();
}, 50);

// Animate tree count
document.addEventListener('DOMContentLoaded', function () {
  const counterEl = document.getElementById('tree-count');
  if (!counterEl) return;

  const target = parseInt(counterEl.getAttribute('data-target'), 10) || 0;

  function animateCounter(el, target, duration = 1600) {
    const start = performance.now();
    const easeOutCubic = t => 1 - Math.pow(1 - t, 3);

    function frame(now) {
      const progress = Math.min((now - start) / duration, 1);
      const eased = easeOutCubic(progress);
      const value = Math.floor(target * eased);
      el.textContent = value.toLocaleString();

      if (progress < 1) {
        requestAnimationFrame(frame);
      } else {
        el.textContent = target.toLocaleString(); // exact final number
      }
    }
    requestAnimationFrame(frame);
  }

  counterEl.textContent = '0'; // reset to 0
  animateCounter(counterEl, target, 1600);
});

// Goal-2026 counter: counts in steps (e.g., 100k) and shows 1M+ badge at >=1,000,000
(function () {
  function formatNumber(n) { return n.toLocaleString(); }

  document.addEventListener('DOMContentLoaded', function () {
    const el = document.getElementById('goal-count');
    const badge = document.getElementById('goal-badge');
    if (!el) return;

    const target = parseInt(el.getAttribute('data-target') || '1000000', 10);
    const step = parseInt(el.getAttribute('data-step') || '100000', 10);
    let current = parseInt(el.getAttribute('data-start') || el.textContent.replace(/,/g,'') || '0', 10);

    // If you want "compact" badge text for targets other than 1,000,000, we compute:
    const compact = new Intl.NumberFormat('en', { notation: 'compact', compactDisplay: 'short' }).format(target) + '+';

    // Reset display to starting point:
    el.textContent = formatNumber(current);
    badge.textContent = compact;

    // If target <= current, show final state immediately
    if (current >= target) {
      el.textContent = formatNumber(target);
      badge.classList.add('visible');
      return;
    }

    // Step-timer: jump in increments of `step` until we hit/meet the target
    const intervalMs = 420; // milliseconds between jumps
    const timer = setInterval(() => {
      current = Math.min(current + step, target);
      el.textContent = formatNumber(current);

      // when we reach >= 1,000,000 show badge "1M+" (or compact computed)
      if (current >= 1000000) {
        badge.classList.add('visible');
      }

      if (current >= target) {
        clearInterval(timer);
        // ensure final text:
        el.textContent = formatNumber(target);
      }
    }, intervalMs);
  });
})();

// Goal-2026 counter with progress bar
(function () {
  function formatNumber(n) { return n.toLocaleString(); }

  document.addEventListener('DOMContentLoaded', function () {
    const el = document.getElementById('goal-count');
    const badge = document.getElementById('goal-badge');
    const bar = document.getElementById('goal-progress-bar');
    const progressText = document.getElementById('goal-progress-text');
    if (!el) return;

    const target = parseInt(el.getAttribute('data-target') || '1000000', 10);
    const step = parseInt(el.getAttribute('data-step') || '100000', 10);
    let current = parseInt(el.getAttribute('data-start') || el.textContent.replace(/,/g,'') || '0', 10);

    const compact = new Intl.NumberFormat('en', { notation: 'compact', compactDisplay: 'short' }).format(target) + '+';

    el.textContent = formatNumber(current);
    badge.textContent = compact;

    function updateProgress(val) {
      const percent = Math.min((val / target) * 100, 100);
      bar.style.width = percent + "%";
      progressText.textContent = Math.floor(percent) + "% towards goal";
    }

    // initialize
    updateProgress(current);

    if (current >= target) {
      el.textContent = formatNumber(target);
      badge.classList.add('visible');
      updateProgress(target);
      return;
    }

    const intervalMs = 420;
    const timer = setInterval(() => {
      current = Math.min(current + step, target);
      el.textContent = formatNumber(current);
      updateProgress(current);

      if (current >= 1000000) {
        badge.classList.add('visible');
      }
      if (current >= target) {
        clearInterval(timer);
        el.textContent = formatNumber(target);
        updateProgress(target);
      }
    }, intervalMs);
  });
})();

function updateProgress(val) {
  const percent = Math.min((val / target) * 100, 100);
  bar.style.width = percent + "%";
  progressText.textContent = Math.floor(percent) + "% towards goal";

  // dynamic colors
  if (percent < 30) {
    bar.style.background = "linear-gradient(90deg, #d32f2f, #f44336)"; // red
  } else if (percent < 70) {
    bar.style.background = "linear-gradient(90deg, #fbc02d, #fdd835)"; // yellow
  } else {
    bar.style.background = "linear-gradient(90deg, #43a047, #2e7d32)"; // green
  }
}

// Generic animated counter with dynamic progress
function initGoalCounter(id, target, step) {
    const el = document.getElementById(`goal-count-${id}`);
    const badge = document.getElementById(`goal-badge-${id}`);
    const bar = document.getElementById(`goal-progress-bar-${id}`);
    const progressText = document.getElementById(`goal-progress-text-${id}`);
    if (!el) return;

    let current = parseInt(el.getAttribute('data-start') || "0", 10);
    el.textContent = current.toLocaleString();

    // Compact badge (e.g., 20B+)
    const compact = new Intl.NumberFormat('en', {
        notation: 'compact',
        compactDisplay: 'short'
    }).format(target) + "+";
    badge.textContent = compact;

    function updateProgress(val) {
        const percent = Math.min((val / target) * 100, 100);
        bar.style.width = percent + "%";
        progressText.textContent = Math.floor(percent) + "% towards goal";

        // dynamic color
        if (percent < 30) {
            bar.style.background = "linear-gradient(90deg, #d32f2f, #f44336)"; // red
            progressText.style.color = "#d32f2f";
        } else if (percent < 70) {
            bar.style.background = "linear-gradient(90deg, #fbc02d, #fdd835)"; // yellow
            progressText.style.color = "#fbc02d";
        } else {
            bar.style.background = "linear-gradient(90deg, #66bb6a, #43a047, #2e7d32)"; // green
        }
    } // Missing closing brace for updateProgress()
} // Missing closing brace for initGoalCounter()