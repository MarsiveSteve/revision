<!DOCTYPE html><html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>LawnCare EA — Lawn Mowing & Tree Planting</title>
  <style>
    :root{
      --brand:#28a745;           /* primary green */
      --brand-2:#5cd85c;         /* accent green */
      --ink:#1f2937;             /* text */
      --muted:#6b7280;           /* secondary text */
      --bg:#f6f8fb;              /* page bg */
      --card:#ffffff;            /* card bg */
      --ring:rgba(40,167,69,.25);
      --shadow:0 10px 25px rgba(0,0,0,.08);
      --radius:18px;
    }
    *{box-sizing:border-box}
    html{scroll-behavior:smooth}
    body{margin:0;font-family:system-ui,-apple-system,Segoe UI,Roboto,Arial,Helvetica,sans-serif;color:var(--ink);background:var(--bg)}/* ---------- NAV ---------- */
.nav{position:sticky;top:0;z-index:50;background:#ffffffcc;backdrop-filter:blur(8px);border-bottom:1px solid #e5e7eb}
.nav .row{max-width:1100px;margin:auto;display:flex;align-items:center;justify-content:space-between;padding:12px 16px}
.brand{display:flex;align-items:center;gap:10px;font-weight:800;color:var(--ink);text-decoration:none}
.brand .dot{width:12px;height:12px;border-radius:999px;background:var(--brand);box-shadow:0 0 0 6px var(--ring)}
.nav a{color:var(--ink);text-decoration:none;font-weight:600;margin-left:18px}
.nav a:hover{color:var(--brand)}

/* ---------- HERO ---------- */
.hero{background:linear-gradient(135deg,var(--brand),#1d9a3d 45%, #128a33);color:#fff}
.hero .wrap{max-width:1100px;margin:auto;padding:80px 16px;display:grid;grid-template-columns:1.1fr .9fr;gap:28px}
.hero h1{font-size:clamp(28px,5vw,46px);line-height:1.1;margin:0 0 16px}
.hero p{font-size:clamp(16px,2.2vw,18px);opacity:.95;margin:0 0 26px}
.cta{display:flex;gap:12px;flex-wrap:wrap}
.btn{appearance:none;border:none;border-radius:999px;padding:12px 18px;font-weight:700;cursor:pointer;text-decoration:none;display:inline-flex;align-items:center;gap:8px}
.btn.primary{background:#fff;color:#0f5132}
.btn.secondary{background:rgba(255,255,255,.14);color:#fff;border:1px solid rgba(255,255,255,.35)}
.hero-card{background:rgba(255,255,255,.12);border:1px solid rgba(255,255,255,.25);border-radius:var(--radius);padding:16px}
.hero ul{margin:0;padding:0 0 0 18px}

/* ---------- SECTION WRAPPER ---------- */
.section{max-width:1100px;margin:60px auto;padding:0 16px}
.section h2{font-size:clamp(22px,3.8vw,32px);margin:0 0 10px}
.muted{color:var(--muted)}

/* ---------- SERVICES ---------- */
.grid{display:grid;gap:16px}
@media (min-width:720px){.grid.cols-3{grid-template-columns:repeat(3,1fr)}}
.card{background:var(--card);border-radius:var(--radius);box-shadow:var(--shadow);padding:18px;border:1px solid #edf0f4}
.card h3{margin:2px 0 8px;font-size:20px}
.price{font-size:28px;font-weight:800;color:var(--brand)}
.chip{display:inline-flex;align-items:center;gap:6px;border-radius:999px;padding:6px 10px;background:#f1f5f9;font-size:12px;margin-right:6px}
.card ul{margin:10px 0 0;padding:0 0 0 18px;color:#4b5563}

/* ---------- BOOKING ---------- */
form{display:grid;gap:14px}
label{font-weight:650}
input, select, textarea{width:100%;padding:12px;border:1px solid #dfe4ea;border-radius:10px;background:#fff;outline:none}
input:focus, select:focus, textarea:focus{border-color:var(--brand);box-shadow:0 0 0 4px var(--ring)}
.form-row{display:grid;gap:12px}
@media (min-width:720px){.form-row.cols-2{grid-template-columns:1fr 1fr}}
.pay-methods{display:flex;gap:10px;flex-wrap:wrap}
.pay-badge{display:inline-flex;align-items:center;gap:8px;border:1px dashed #c9d4e5;padding:8px 12px;border-radius:999px;background:#fff}
.summary{background:#f8fafc;border:1px solid #e5edf7;border-radius:14px;padding:14px}

/* ---------- TESTIMONIALS / CTA STRIP ---------- */
.strip{background:#0f5132;color:#eafff1;border-radius:var(--radius);padding:22px;display:flex;align-items:center;justify-content:space-between;gap:12px}
.strip .btn{background:#eafff1;color:#0f5132}

/* ---------- FOOTER ---------- */
footer{margin-top:60px;padding:30px 16px;color:#64748b;background:#0b1d13}
footer .foot{max-width:1100px;margin:auto;display:grid;gap:16px}
footer a{color:#a6ffcb;text-decoration:none}

/* ---------- MODAL (Demo only) ---------- */
.modal-backdrop{position:fixed;inset:0;background:rgba(0,0,0,.5);display:none;align-items:center;justify-content:center;padding:16px}
.modal{background:#fff;border-radius:16px;max-width:520px;width:100%;padding:18px;box-shadow:var(--shadow)}
.modal h3{margin:0 0 8px}
.modal .actions{display:flex;justify-content:flex-end;gap:10px;margin-top:12px}

  </style>
</head>
<body>
  <!-- NAV -->
  <nav class="nav">
    <div class="row">
      <a href="#home" class="brand"><span class="dot"></span> LawnCare EA</a>
      <div>
        <a href="#services">Services</a>
        <a href="#booking">Booking</a>
        <a href="#about">About</a>
        <a href="#contact">Contact</a>
      </div>
    </div>
  </nav>  <!-- HERO -->  <header id="home" class="hero">
    <div class="wrap">
      <div>
        <h1>Professional Lawn Care & Tree Planting Across East Africa</h1>
        <p>Transform your outdoor spaces while contributing to environmental conservation. Expert lawn mowing services and community tree planting initiatives. Display‑only build — ready for payment/API wiring.</p>
        <div class="cta">
          <a class="btn primary" href="#booking">Book Service</a>
          <a class="btn secondary" href="#services">View Pricing</a>
        </div>
      </div>
      <div class="hero-card">
        <strong>Why choose us?</strong>
        <ul>
          <li>Reliable, vetted crews in major EA cities</li>
          <li>Eco‑friendly tools & water‑wise practices</li>
          <li>Transparent KSh pricing, no surprises</li>
        </ul>
      </div>
    </div>
  </header>  <!-- SERVICES -->  <section id="services" class="section">
    <h2>Services & Pricing</h2>
    <p class="muted">Simple, transparent packages. Prices in Kenyan Shillings (KSh).</p>
    <div class="grid cols-3">
      <article class="card">
        <span class="chip">Small Yard</span>
        <h3>Lawn Mowing — Small</h3>
        <div class="price" data-price="1500">KSh 1,500</div>
        <ul>
          <li>Up to ~250 m²</li>
          <li>Trim edges & quick blow‑off</li>
          <li>Bag or mulch clippings</li>
        </ul>
      </article>
      <article class="card">
        <span class="chip">Medium Yard</span>
        <h3>Lawn Mowing — Medium</h3>
        <div class="price" data-price="2500">KSh 2,500</div>
        <ul>
          <li>~250–600 m²</li>
          <li>Edge trimming included</li>
          <li>Clippings removal</li>
        </ul>
      </article>
      <article class="card">
        <span class="chip">Large Yard</span>
        <h3>Lawn Mowing — Large</h3>
        <div class="price" data-price="4000">KSh 4,000</div>
        <ul>
          <li>600+ m²</li>
          <li>Full perimeter edging</li>
          <li>Debris & green‑waste tidy</li>
        </ul>
      </article>
    </div><div class="grid" style="margin-top:16px">
  <article class="card">
    <h3>Add‑ons</h3>
    <ul>
      <li>Hedge trimming <strong>(KSh 800)</strong></li>
      <li>Tree planting per sapling <strong>(KSh 300)</strong></li>
      <li>Green‑waste removal <strong>(KSh 500)</strong></li>
    </ul>
  </article>
</div>

  </section>  <!-- BOOKING -->  <section id="booking" class="section">
    <h2>Book a Service</h2>
    <p class="muted">Display‑only. Form totals update instantly. Hook your database & payments later.</p><div class="grid" style="grid-template-columns:2fr 1.2fr; gap:16px">
  <!-- Form -->
  <article class="card">
    <form id="bookingForm">
      <div class="form-row cols-2">
        <div>
          <label for="fullName">Full name</label>
          <input id="fullName" name="fullName" placeholder="Jane Doe" required />
        </div>
        <div>
          <label for="phone">Phone (M‑Pesa/Airtel Money)</label>
          <input id="phone" name="phone" inputmode="tel" placeholder="07xx xxx xxx" required />
        </div>
      </div>

      <div class="form-row cols-2">
        <div>
          <label for="email">Email</label>
          <input id="email" name="email" type="email" placeholder="jane@example.com" />
        </div>
        <div>
          <label for="date">Preferred date</label>
          <input id="date" name="date" type="date" required />
        </div>
      </div>

      <div>
        <label for="address">Service address</label>
        <input id="address" name="address" placeholder="Street, Estate/Area, City" required />
      </div>

      <div class="form-row cols-2">
        <div>
          <label for="size">Lawn size</label>
          <select id="size" name="size" required>
            <option value="small" data-price="1500">Small — KSh 1,500</option>
            <option value="medium" data-price="2500">Medium — KSh 2,500</option>
            <option value="large" data-price="4000">Large — KSh 4,000</option>
          </select>
        </div>
        <div>
          <label for="frequency">Frequency</label>
          <select id="frequency" name="frequency">
            <option value="once" data-mult="1">One‑time</option>
            <option value="biweekly" data-mult="0.95">Bi‑weekly (‑5%)</option>
            <option value="monthly" data-mult="0.98">Monthly (‑2%)</option>
          </select>
        </div>
      </div>

      <div class="form-row cols-2">
        <div>
          <label><input type="checkbox" id="addonHedge" data-fee="800" /> Hedge trimming (KSh 800)</label>
        </div>
        <div>
          <label><input type="checkbox" id="addonWaste" data-fee="500" /> Green‑waste removal (KSh 500)</label>
        </div>
      </div>

      <div>
        <label for="saplings">Tree planting — number of saplings</label>
        <input id="saplings" type="number" min="0" value="0" />
        <div class="muted" style="font-size:12px">KSh 300 per sapling</div>
      </div>

      <div>
        <label>Payment method (UI only)</label>
        <div class="pay-methods">
          <span class="pay-badge" data-method="mpesa">🟢 M‑Pesa</span>
          <span class="pay-badge" data-method="airtel">🔴 Airtel Money</span>
          <span class="pay-badge" data-method="paypal">💙 PayPal</span>
        </div>
      </div>

      <div class="summary" id="summaryBox">
        <strong>Order Summary</strong>
        <div id="summaryLines" class="muted" style="margin-top:6px"></div>
        <div style="display:flex;justify-content:space-between;align-items:center;margin-top:10px">
          <span>Total</span>
          <strong id="totalDisplay" style="font-size:22px;color:var(--brand)">KSh 0</strong>
        </div>
      </div>

      <div style="display:flex;gap:10px;flex-wrap:wrap">
        <button type="submit" class="btn" style="background:var(--brand);color:#fff">Confirm Booking (Demo)</button>
        <a href="#services" class="btn" style="background:#eef6f0;color:#0f5132">Back to Pricing</a>
      </div>

      <!-- Integration notes (placeholders) -->
      <small class="muted">
        Integration hooks (to wire later): <code>window.payments.mpesa.stkPush()</code>, <code>window.payments.airtel.charge()</code>, <code>window.payments.paypal.checkout()</code> — see JS TODOs below.
      </small>
    </form>
  </article>

  <!-- Help/Assurance -->
  <aside class="card">
    <h3>What happens next?</h3>
    <p class="muted">You’ll receive a confirmation SMS/Email. A crew lead will call to confirm access and any special notes. For recurring plans, we’ll keep the same day/time slot.</p>
    <div class="strip" style="margin-top:10px">
      <div>
        <strong>Plant a tree with every mow</strong><br>
        <span class="muted" style="color:#d2ffe6">Add saplings in your booking — we’ll plant them for you.</span>
      </div>
      <a href="#booking" class="btn">Add Saplings</a>
    </div>
  </aside>
</div>

  </section>  <!-- ABOUT -->  <section id="about" class="section">
    <h2>About LawnCare EA</h2>
    <p class="muted">We’re a locally‑owned team operating across Nairobi, Mombasa, Kisumu, Kampala & Dar es Salaam. Our mission is simple: gorgeous lawns, happy clients, and greener neighborhoods.</p>
    <div class="grid cols-3" style="margin-top:12px">
      <article class="card">
        <h3>Eco Practices</h3>
        <p>Electric mowers where viable, water‑wise scheduling, and responsible green‑waste disposal.</p>
      </article>
      <article class="card">
        <h3>Trusted Crews</h3>
        <p>Background‑checked, trained teams that arrive on time with the right tools for your lawn size.</p>
      </article>
      <article class="card">
        <h3>Community Impact</h3>
        <p>We partner with local groups to plant indigenous trees and restore urban biodiversity.</p>
      </article>
    </div>
  </section>  <!-- CONTACT -->  <section id="contact" class="section">
    <h2>Contact</h2>
    <div class="grid" style="grid-template-columns:1.3fr 1fr;gap:16px">
      <article class="card">
        <h3>Message us</h3>
        <form onsubmit="event.preventDefault(); openModal('Thanks! We\'ll get back to you shortly. (Demo)')">
          <div class="form-row cols-2">
            <input placeholder="Your name" required />
            <input type="email" placeholder="Email" required />
          </div>
          <textarea rows="4" placeholder="How can we help?" required></textarea>
          <button class="btn" style="background:var(--brand);color:#fff">Send Message</button>
        </form>
      </article>
      <aside class="card">
        <h3>Reach us</h3>
        <p class="muted">Mon–Sat 8:00–18:00 EAT</p>
        <p>📞 +254 700 000 000<br>✉️ hello@lawncare‑ea.example</p>
        <p class="muted">Nairobi • Mombasa • Kisumu • Kampala • Dar es Salaam</p>
      </aside>
    </div>
  </section>  <!-- FOOTER -->  <footer>
    <div class="foot">
      <div>&copy; <span id="year"></span> LawnCare EA. All rights reserved.</div>
      <div><a href="#home">Back to top ↑</a></div>
    </div>
  </footer>  <!-- DEMO MODAL -->  <div class="modal-backdrop" id="modal">
    <div class="modal">
      <h3 id="modalTitle">Booking received</h3>
      <p id="modalBody" class="muted">This is a display‑only demo. No payment was processed.</p>
      <div class="actions">
        <button class="btn" onclick="closeModal()" style="background:#eef2ff">Close</button>
      </div>
    </div>
  </div>  <script>
    // ---------- Utils ----------
    const fmtKES = n => `KSh ${n.toLocaleString('en-KE')}`;

    // ---------- Dynamic year ----------
    document.getElementById('year').textContent = new Date().getFullYear();

    // ---------- Booking summary & totals ----------
    const sizeSel = document.getElementById('size');
    const freqSel = document.getElementById('frequency');
    const addonHedge = document.getElementById('addonHedge');
    const addonWaste = document.getElementById('addonWaste');
    const saplings = document.getElementById('saplings');
    const linesBox = document.getElementById('summaryLines');
    const totalDisplay = document.getElementById('totalDisplay');

    const PRICES = { small:1500, medium:2500, large:4000, hedge:800, waste:500, sapling:300 };

    function recalc(){
      const size = sizeSel.value;
      const base = PRICES[size];
      const freqMult = Number(freqSel.options[freqSel.selectedIndex].dataset.mult);
      const hedge = addonHedge.checked ? PRICES.hedge : 0;
      const waste = addonWaste.checked ? PRICES.waste : 0;
      const saplingCount = Math.max(0, Number(saplings.value||0));
      const saplingFee = saplingCount * PRICES.sapling;

      const subtotal = base + hedge + waste + saplingFee;
      const total = Math.round(subtotal * freqMult);

      // Render lines
      const parts = [
        `Base (${size.charAt(0).toUpperCase()+size.slice(1)}) — ${fmtKES(base)}`,
        hedge? `Hedge trimming — ${fmtKES(hedge)}`: null,
        waste? `Green‑waste removal — ${fmtKES(waste)}`: null,
        saplingCount? `Tree planting x${saplingCount} — ${fmtKES(saplingFee)}`: null,
        freqMult !== 1 ? `Frequency discount applied` : null
      ].filter(Boolean);
      linesBox.innerHTML = parts.map(p=>`<div>• ${p}</div>`).join('');
      totalDisplay.textContent = fmtKES(total);
    }

    [sizeSel, freqSel, addonHedge, addonWaste, saplings].forEach(el=>{
      el.addEventListener('input', recalc);
      el.addEventListener('change', recalc);
    });
    recalc();

    // ---------- Form submit (Demo) ----------
    const form = document.getElementById('bookingForm');
    form.addEventListener('submit', (e)=>{
      e.preventDefault();
      openModal("This is a demo submission. Wire your database + payments to go live.");
    });

    // ---------- Modal helpers ----------
    function openModal(msg){
      document.getElementById('modalBody').textContent = msg;
      document.getElementById('modal').style.display = 'flex';
    }
    function closeModal(){
      document.getElementById('modal').style.display = 'none';
    }

    // ---------- Payment Integration Hooks (TODOs) ----------
    // Wire these up when you add your backend:
    // window.payments = {
    //   mpesa: {
    //     // Example: call your server to initiate STK Push (Safaricom Daraja API)
    //     // async stkPush(msisdn, amount, reference){
    //     //   return fetch('/api/pay/mpesa', {method:'POST', headers:{'Content-Type':'application/json'}, body:JSON.stringify({ msisdn, amount, reference })});
    //     // }
    //   },
    //   airtel: {
    //     // async charge(msisdn, amount){
    //     //   return fetch('/api/pay/airtel', {method:'POST', headers:{'Content-Type':'application/json'}, body:JSON.stringify({ msisdn, amount })});
    //     // }
    //   },
    //   paypal: {
    //     // The typical flow uses PayPal JS SDK; render a button and create an order via your server
    //     // See: https://developer.paypal.com/docs/checkout/ 
    //   }
    // };
  </script></body>
</html>