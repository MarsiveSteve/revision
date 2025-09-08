<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Marsive GreenCut EA</title>
  <style>
    body { margin: 0; font-family: Arial, sans-serif; line-height: 1.6; }
    header {
      background: #2e7d32; color: #fff; padding: 15px 20px;
      display: flex; justify-content: space-between; align-items: center;
      position: sticky; top: 0; z-index: 1000;
    }
    header h1 { font-size: 20px; display: flex; align-items: center; gap: 8px; }
    nav a { color: white; margin: 0 10px; text-decoration: none; font-weight: bold; }
    nav a:hover { text-decoration: underline; }
    .hero {
      background: linear-gradient(to right, #43a047, #66bb6a); color: white;
      text-align: center; padding: 60px 20px;
    }
    .hero h2 { font-size: 32px; margin-bottom: 20px; }
    section { padding: 50px 20px; max-width: 1100px; margin: auto; }
    section h2 { text-align: center; margin-bottom: 30px; font-size: 28px; color: #2e7d32; }
    .grid {
      display: grid; grid-template-columns: repeat(auto-fit, minmax(160px, 1fr)); gap: 15px;
    }
    .card {
      background: #f9f9f9; padding: 15px; border-radius: 10px;
      box-shadow: 0 2px 5px rgba(0,0,0,0.1); text-align: center;
      font-size: 14px; font-weight: bold; color: #2e7d32;
      cursor: pointer; transition: transform 0.2s;
    }
    .card:hover { transform: scale(1.05); background: #e8f5e9; }
    .modal {
      display: none; position: fixed; z-index: 2000; left: 0; top: 0;
      width: 100%; height: 100%; background: rgba(0,0,0,0.6);
      display: flex; align-items: center; justify-content: center;
    }
    .modal-content {
      background: #fff; padding: 30px; border-radius: 12px;
      max-width: 400px; text-align: center; box-shadow: 0 4px 10px rgba(0,0,0,0.3);
    }
    .modal-content h3 { color: #2e7d32; }
    .modal-content img {
      max-width: 80px; margin: 15px auto; display: block;
    }
    .close-btn {
      margin-top: 20px; padding: 10px 20px; background: #2e7d32; color: white;
      border: none; border-radius: 6px; cursor: pointer;
    }
    .close-btn:hover { background: #1b5e20; }
    footer { background: #2e7d32; color: white; text-align: center; padding: 20px; margin-top: 40px; }
    
    /* searchable dropdown */
    #regionList {
      max-height: 180px; overflow-y: auto; border: 1px solid #ccc; 
      border-radius: 6px; width: 260px; margin: 10px auto; 
      background: #fff; text-align: left; display: none;
    }
    #regionList .region-item { padding: 8px 12px; cursor: pointer; }
    #regionList .region-item:hover { background: #e8f5e9; }
  </style>
</head>
<body>
  <header>
    <h1>🌿 Marsive GreenCut EA</h1>
    <nav>
      <a href="#counties">Counties</a>
      <a href="#ea">East Africa</a>
      <a href="#contact">Contact</a>
    </nav>
  </header>

  <div class="hero">
    <h2>Professional Lawn Care & Tree Planting Across East Africa</h2>
    <p>Transforming outdoor spaces while contributing to environmental conservation.</p>
  </div>

  <!-- Counties -->
  <section id="counties">
    <h2>We Serve All 47 Counties in Kenya</h2>

    <!-- Searchable Selector -->
    <section id="quicknav" style="text-align:center; margin:20px 0;">
      <input 
        type="text" 
        id="searchBox" 
        placeholder="Search county or country..." 
        style="padding:10px; font-size:16px; border-radius:6px; border:1px solid #ccc; min-width:250px;"
        onkeyup="filterRegions()"
      >
      <div id="regionList">
        <div class="region-item" onclick="selectRegion('Nairobi')">Nairobi</div>
        <div class="region-item" onclick="selectRegion('Mombasa')">Mombasa</div>
        <div class="region-item" onclick="selectRegion('Kisumu')">Kisumu</div>
        <div class="region-item" onclick="selectRegion('Nakuru')">Nakuru</div>
        <div class="region-item" onclick="selectRegion('Kiambu')">Kiambu</div>
        <div class="region-item" onclick="selectRegion('Tanzania')">🇹🇿 Tanzania</div>
        <div class="region-item" onclick="selectRegion('Uganda')">🇺🇬 Uganda</div>
        <div class="region-item" onclick="selectRegion('Rwanda')">🇷🇼 Rwanda</div>
        <div class="region-item" onclick="selectRegion('Burundi')">🇧🇮 Burundi</div>
      </div>
    </section>

    <!-- County Grid -->
    <div class="grid">
      <div class="card" onclick="showModal('Nairobi')">Nairobi</div>
      <div class="card" onclick="showModal('Mombasa')">Mombasa</div>
      <div class="card" onclick="showModal('Kisumu')">Kisumu</div>
      <div class="card" onclick="showModal('Nakuru')">Nakuru</div>
      <div class="card" onclick="showModal('Kiambu')">Kiambu</div>
    </div>
  </section>

  <!-- East Africa -->
  <section id="ea">
    <h2>Expanding Across East Africa</h2>
    <div class="grid">
      <div class="card" onclick="showModal('Tanzania')">🇹🇿 Tanzania</div>
      <div class="card" onclick="showModal('Uganda')">🇺🇬 Uganda</div>
      <div class="card" onclick="showModal('Rwanda')">🇷🇼 Rwanda</div>
      <div class="card" onclick="showModal('Burundi')">🇧🇮 Burundi</div>
    </div>
  </section>

  <!-- Modal -->
  <div id="countyModal" class="modal">
    <div class="modal-content">
      <img id="regionIcon" src="" alt="">
      <h3 id="countyName"></h3>
      <p id="countyInfo"></p>
      <button class="close-btn" onclick="closeModal()">Close</button>
    </div>
  </div>

  <footer>
    <p>🌿 Marsive GreenCut EA &copy; 2025</p>
  </footer>

  <script>
    const messages = {
      "Nairobi": { text: "Corporate packages available for offices, estates & rooftop gardens.", icon: "https://img.icons8.com/emoji/96/deciduous-tree.png" },
      "Mombasa": { text: "Special coastal lawn care services with drought-resistant options.", icon: "https://img.icons8.com/emoji/96/palm-tree.png" },
      "Kisumu": { text: "Community tree planting hub around Lake Victoria.", icon: "https://img.icons8.com/emoji/96/seedling.png" },
      "Nakuru": { text: "Eco-friendly estate and gated community projects.", icon: "https://img.icons8.com/emoji/96/herb.png" },
      "Kiambu": { text: "Premium lawn care for residential and golf courses.", icon: "https://img.icons8.com/emoji/96/evergreen-tree.png" },
      "Tanzania": { text: "Expanding to Dar & Arusha with corporate services.", icon: "https://img.icons8.com/emoji/96/flag-tanzania.png" },
      "Uganda": { text: "Launching Kampala green schools program.", icon: "https://img.icons8.com/emoji/96/flag-uganda.png" },
      "Rwanda": { text: "Eco-friendly Kigali greening initiatives.", icon: "https://img.icons8.com/emoji/96/flag-rwanda.png" },
      "Burundi": { text: "Community agroforestry & lawns in Bujumbura.", icon: "https://img.icons8.com/emoji/96/flag-burundi.png" }
    };

    function showModal(region) {
      const info = messages[region] || { text: "We proudly serve " + region + ".", icon: "https://img.icons8.com/emoji/96/seedling.png" };
      document.getElementById("countyName").textContent = region;
      document.getElementById("countyInfo").textContent = info.text;
      document.getElementById("regionIcon").src = info.icon;
      document.getElementById("countyModal").style.display = "flex";
    }

    function closeModal() {
      document.getElementById("countyModal").style.display = "none";
    }

    function filterRegions() {
      const input = document.getElementById("searchBox");
      const filter = input.value.toLowerCase();
      const list = document.getElementById("regionList");
      const items = list.getElementsByClassName("region-item");

      list.style.display = "block"; // show list while typing

      for (let i = 0; i < items.length; i++) {
        let txt = items[i].textContent || items[i].innerText;
        items[i].style.display = txt.toLowerCase().indexOf(filter) > -1 ? "" : "none";
      }
    }

    function selectRegion(region) {
      document.getElementById("searchBox").value = region;
      document.getElementById("regionList").style.display = "none";
      showModal(region);
    }
  </script>
</body>
</html>