<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Marsive GreenCut EA</title>
  <style>
    :root { 
      --green:#2e7d32; 
      --light:#f4f4f4; 
    }

    body { 
      margin:0; 
      font-family:Arial, Helvetica, sans-serif; 
      background:var(--light); 
      color:#222; 
      line-height:1.6;
    }

    /* Header */
    header { 
      background:var(--green); 
      color:#fff; 
      padding:16px 40px; 
      position:sticky; 
      top:0; 
      z-index:100; 
      display:flex; 
      justify-content:space-between; 
      align-items:center; 
    }
    header h1 { 
      margin:0; 
      font-size:22px; 
      display:flex; 
      gap:.5rem; 
      align-items:center; 
    }
    nav a { 
      color:#fff; 
      text-decoration:none; 
      margin-left:20px; 
      font-weight:600; 
      font-size:15px;
    }
    nav a:hover { text-decoration:underline; }

    /* Hero */
    .hero { 
      background: linear-gradient(90deg,#43a047,#66bb6a); 
      color:white; 
      text-align:center; 
      padding:100px 20px;   
    }
    .hero h2 {
      font-size:32px;
      margin-bottom:10px;
    }
    .hero p {
      font-size:18px;
      max-width:700px;
      margin:0 auto;
    }

    /* Section */
    section { 
      padding:60px 40px;   
      max-width:1400px;    
      margin:0 auto; 
    }
    section h2 {
      text-align:center;
      font-size:28px;
      margin-bottom:30px;
      color:var(--green);
    }

    /* Grid */
    .grid { 
      display:grid; 
      grid-template-columns:repeat(auto-fit,minmax(220px,1fr)); 
      gap:20px; 
    }
    .card { 
      background:#fff; 
      border-radius:12px; 
      padding:20px; 
      box-shadow:0 3px 8px rgba(0,0,0,0.1); 
      color:var(--green); 
      font-weight:700; 
      cursor:pointer; 
      text-align:center; 
      transition:all .2s ease;
    }
    .card:hover { 
      transform:translateY(-3px); 
      box-shadow:0 6px 16px rgba(0,0,0,0.15); 
    }

    /* Mission & Vision */
    .mv-section {
      background:white;
      border-radius:12px;
      padding:40px;
      box-shadow:0 3px 8px rgba(0,0,0,0.1);
      max-width:900px;
      margin:0 auto;
    }
    .mv-block {
      margin-bottom:30px;
    }
    .mv-block h3 {
      color:var(--green);
      margin-bottom:10px;
      font-size:22px;
    }

    /* Contact */
    .contact { 
      text-align:center; 
    }
    .contact p {
      margin:8px 0;
    }
    .btn {
      background:var(--green);
      color:white;
      padding:12px 24px;
      border:none;
      border-radius:6px;
      cursor:pointer;
      font-size:16px;
      font-weight:bold;
      margin-top:12px;
      transition:.2s;
    }
    .btn:hover { background:#1b5e20; }

    /* Footer */
    footer {
      background:#222;
      color:#ddd;
      text-align:center;
      padding:20px;
      font-size:14px;
      margin-top:40px;
    }
    
    /* Call-to-Action */
.cta {
  background: linear-gradient(90deg,#2e7d32,#43a047);
  color:white;
  text-align:center;
  padding:80px 20px;
  border-radius:0;
  margin:40px 0;
}
.cta h2 {
  font-size:30px;
  margin-bottom:20px;
}
.cta p {
  font-size:18px;
  max-width:700px;
  margin:0 auto 25px;
}
.cta .btn {
  background:#fff;
  color:var(--green);
  font-size:18px;
  padding:14px 28px;
  border-radius:8px;
}
.cta .btn:hover {
  background:#f4f4f4;
}
  </style>
</head>
<body>

  <!-- Header -->
  <header>
    <h1>Marsive GreenCut EA</h1>
    <nav>
      <a href="#about">About</a>
      <a href="#mission">Mission & Vision</a>
      <a href="#services">Services</a>
      <a href="#contact">Contact</a>
    </nav>
  </header>

  <!-- Hero -->
  <div class="hero">
    <h2>Professional Lawn Care & Tree Planting Across East Africa</h2>
    <p>Transforming outdoor spaces while contributing to environmental conservation.</p>
  </div>

  <!-- About -->
  <section id="about">
    <h2>About Us</h2>
    <p style="text-align:center; max-width:800px; margin:0 auto;">
      Marsive GreenCut EA provides professional lawn care, landscaping, and 
      tree planting services across all 47 counties in Kenya and beyond. 
      Our mission is to beautify outdoor spaces while promoting sustainable 
      environmental practices.
    </p>
  </section>

  <!-- Mission & Vision -->
  <section id="mission">
    <h2>Mission & Vision</h2>
    <div class="mv-section">
      <div class="mv-block">
        <h3>🌍 Mission</h3>
        <p>
          To have functional, climate-related activities in each county regardless of the climate, ensuring sustainable growth and environmental impact at every level.
        </p>
      </div>
      <div class="mv-block">
        <h3>🌱 Vision</h3>
        <p>
          To be the biggest player in climate change efforts by planting 20+ billion trees in 10 years, with the help of everyone across the globe, making the planet better than we found it.
        </p>
      </div>
    </div>
  </section>

  <!-- Services -->
  <section id="services">
    <h2>Our Services</h2>
    <div class="grid">
      <div class="card">🌱 Lawn Care</div>
      <div class="card">🌳 Tree Planting</div>
      <div class="card">🌿 Landscaping</div>
      <div class="card">🚜 Garden Maintenance</div>
      <div class="card">💧 Irrigation Setup</div>
      <div class="card">🌾 Environmental Conservation</div>
    </div>
  </section>

  <!-- Contact -->
  <section id="contact">
    <h2>Contact Us</h2>
    <div class="contact">
      <p>Email: info@marsivegreencut.com</p>
      <p>Phone: +254 700 123 456</p>
      <p>Serving all 47 counties in Kenya & East Africa</p>
      <button class="btn">Book a Service</button>
    </div>
  </section>
  
  <!-- Call to Action -->
<section class="cta">
  <div class="cta-box">
    <h2>🌍 Be Part of the Change</h2>
    <p>
      Together, we can plant 20+ billion trees in the next 10 years. 
      Your support will help make the planet greener and better than we found it.
    </p>
    <button class="btn">Join the Movement</button>
  </div>
</section>

  <!-- Footer -->
  <footer>
    <p>&copy; 2025 Marsive GreenCut EA. All rights reserved.</p>
  </footer>

</body>
</html>