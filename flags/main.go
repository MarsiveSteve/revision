<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Marsive GreenCut East Africa</title>

  <!-- Bootstrap CSS -->
  <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet">

  <!-- FontAwesome -->
  <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.0/css/all.min.css" rel="stylesheet">

  <!-- Google Fonts -->
  <link href="https://fonts.googleapis.com/css2?family=Poppins:wght@400;600&display=swap" rel="stylesheet">

  <style>
    body {
      font-family: 'Poppins', sans-serif;
    }

    /* Hero Section */
    .hero {
      background: url("https://picsum.photos/id/1043/1200/600") center/cover no-repeat;
      color: white;
      min-height: 100vh;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: flex-start;
      padding: 50px;
    }
    .hero h1 {
      font-size: clamp(2rem, 5vw, 3.5rem);
      font-weight: 700;
    }
    .hero p {
      font-size: 1.1rem;
      max-width: 600px;
    }
    .btn-cta {
      border: 2px solid white;
      color: white;
      padding: 10px 20px;
      text-decoration: none;
      transition: 0.3s;
    }
    .btn-cta:hover {
      background: white;
      color: #000;
    }

    /* Dark Section */
    .section-dark {
      background: #051431;
      color: white;
      padding: 60px 20px;
    }
    .section-dark img {
      border-radius: 10px;
      max-width: 100%;
    }

    /* Why Choose Us */
    .why-choose {
      padding: 60px 20px;
      background: #f9f9f9;
    }
    .why-choose h2 {
      font-weight: 600;
      margin-bottom: 40px;
    }
    .feature-box {
      background: white;
      padding: 30px;
      border-radius: 12px;
      box-shadow: 0 4px 12px rgba(0,0,0,0.1);
      transition: transform 0.3s;
    }
    .feature-box:hover {
      transform: translateY(-8px);
    }
    .feature-box i {
      font-size: 2rem;
      color: #2e7d32;
      margin-bottom: 15px;
    }

    /* Testimonials */
    .testimonials {
      padding: 60px 20px;
      background: #fff;
    }
    .testimonials h2 {
      font-weight: 600;
      margin-bottom: 40px;
    }
    .testimonial-card {
      background: #f9f9f9;
      padding: 25px;
      border-radius: 12px;
      box-shadow: 0 4px 12px rgba(0,0,0,0.08);
      height: 100%;
      display: flex;
      flex-direction: column;
      justify-content: space-between;
    }
    .testimonial-card p {
      font-style: italic;
    }
    .testimonial-author {
      margin-top: 15px;
      font-weight: 600;
    }
    .stars {
      color: #f4c542;
    }

    /* Booking */
    .booking {
      padding: 60px 20px;
      background: #f9f9f9;
    }
    .booking h2 {
      font-weight: 600;
      margin-bottom: 30px;
    }
    .form-control {
      border-radius: 8px;
      box-shadow: none;
    }
    .btn-submit {
      background: #2e7d32;
      color: white;
      border-radius: 8px;
      padding: 10px 20px;
      transition: 0.3s;
    }
    .btn-submit:hover {
      background: #256428;
    }
  </style>
</head>
<body>

  <!-- Navbar -->
  <nav class="navbar navbar-expand-lg navbar-light" style="background-color:#d4b48c;">
    <div class="container-fluid">
      <a class="navbar-brand fw-bold" href="#">Marsive GreenCut East Africa</a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" 
        aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav ms-auto mb-2 mb-lg-0">
          <li class="nav-item"><a class="nav-link" href="#">Our Services</a></li>
          <li class="nav-item"><a class="nav-link" href="#">Welcome</a></li>
          <li class="nav-item"><a class="nav-link" href="#">Why Choose Us</a></li>
          <li class="nav-item"><a class="nav-link" href="#">Testimonials</a></li>
          <li class="nav-item"><a class="nav-link" href="#">Booking</a></li>
        </ul>
        <a href="#booking" class="btn btn-success ms-lg-3 mt-2 mt-lg-0">Book Service</a>
      </div>
    </div>
  </nav>

  <!-- Hero Section -->
  <div class="hero">
    <h1>Comprehensive Lawn Care <br> Expert Landscaping Solutions</h1>
    <p>Our skilled team specializes in lawn mowing, tree planting, and garden maintenance, ensuring your outdoor space thrives with beauty and sustainability.</p>
    <a href="#booking" class="btn-cta mt-3">Contact Us</a>
  </div>

  <!-- Services Section -->
  <div class="section-dark text-center">
    <h2 class="mb-4">Expert Lawn Care</h2>
    <p class="mb-4">Our professional lawn mowing services ensure that your grass is healthy and beautifully maintained, enhancing the aesthetic of your property.</p>
    <img src="https://picsum.photos/id/1018/1000/500" alt="Lawn Care">
  </div>

  <!-- Why Choose Us Section -->
  <div class="why-choose text-center">
    <div class="container">
      <h2>Why Choose Us</h2>
      <div class="row g-4">
        <div class="col-md-4">
          <div class="feature-box">
            <i class="fa-solid fa-leaf"></i>
            <h5>Eco-Friendly Solutions</h5>
            <p>We use environmentally safe methods to maintain your lawn while protecting nature.</p>
          </div>
        </div>
        <div class="col-md-4">
          <div class="feature-box">
            <i class="fa-solid fa-user-tie"></i>
            <h5>Expert Team</h5>
            <p>Our skilled professionals bring years of experience in landscaping and lawn care.</p>
          </div>
        </div>
        <div class="col-md-4">
          <div class="feature-box">
            <i class="fa-solid fa-clock"></i>
            <h5>Reliable Service</h5>
            <p>We deliver timely, consistent, and high-quality service that you can depend on.</p>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Testimonials Section -->
  <div class="testimonials text-center">
    <div class="container">
      <h2>What Our Clients Say</h2>
      <div class="row g-4">
        <div class="col-md-4">
          <div class="testimonial-card">
            <p>"Marsive GreenCut transformed my garden into a paradise. The team was professional and attentive to detail."</p>
            <div class="stars mb-2">
              <i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-regular fa-star"></i>
            </div>
            <div class="testimonial-author">— Sarah M.</div>
          </div>
        </div>
        <div class="col-md-4">
          <div class="testimonial-card">
            <p>"They did an amazing job with our lawn and even advised us on eco-friendly maintenance tips. Highly recommend!"</p>
            <div class="stars mb-2">
              <i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i>
            </div>
            <div class="testimonial-author">— James K.</div>
          </div>
        </div>
        <div class="col-md-4">
          <div class="testimonial-card">
            <p>"Reliable, affordable, and very professional. I’ll definitely continue using their services."</p>
            <div class="stars mb-2">
              <i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star"></i><i class="fa-solid fa-star-half-stroke"></i><i class="fa-regular fa-star"></i>
            </div>
            <div class="testimonial-author">— Anita L.</div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- Booking Section -->
  <div id="booking" class="booking text-center">
    <div class="container">
      <h2>Book Our Services</h2>
      <p class="mb-4">Fill in the form below and we will get back to you shortly.</p>
      <div class="row justify-content-center">
        <div class="col-md-8">
          <form>
            <div class="mb-3">
              <input type="text" class="form-control" placeholder="Your Name" required>
            </div>
            <div class="mb-3">
              <input type="email" class="form-control" placeholder="Your Email" required>
            </div>
            <div class="mb-3">
              <select class="form-control" required>
                <option value="">Select Service</option>
                <option value="lawn">Lawn Care</option>
                <option value="tree">Tree Planting</option>
                <option value="landscaping">Landscaping</option>
              </select>
            </div>
            <div class="mb-3">
              <textarea class="form-control" rows="4" placeholder="Your Message"></textarea>
            </div>
            <button type="submit" class="btn-submit">Submit Booking</button>
          </form>
        </div>
      </div>
    </div>
  </div>

  <!-- Bootstrap JS -->
  <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js"></script>
</body>
</html>