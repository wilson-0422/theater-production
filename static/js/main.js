document.addEventListener('DOMContentLoaded', function () {
  var deleteForms = document.querySelectorAll('form[onsubmit]');
  deleteForms.forEach(function (form) {
    var originalOnsubmit = form.getAttribute('onsubmit');
    if (originalOnsubmit && originalOnsubmit.indexOf('confirm') !== -1) {
      form.removeAttribute('onsubmit');
      form.addEventListener('submit', function (e) {
        if (!confirm('确认执行此操作？')) {
          e.preventDefault();
        }
      });
    }
  });

  var dateInputs = document.querySelectorAll('input[type="date"]');
  dateInputs.forEach(function (input) {
    if (!input.value) {
      input.value = new Date().toISOString().split('T')[0];
    }
  });

  var datetimeInputs = document.querySelectorAll('input[type="datetime-local"]');
  datetimeInputs.forEach(function (input) {
    if (!input.value) {
      var now = new Date();
      var offset = now.getTimezoneOffset();
      var local = new Date(now.getTime() - offset * 60000);
      input.value = local.toISOString().slice(0, 16);
    }
  });

  var alerts = document.querySelectorAll('.alert');
  alerts.forEach(function (alert) {
    setTimeout(function () {
      alert.style.transition = 'opacity 0.5s';
      alert.style.opacity = '0';
      setTimeout(function () {
        alert.remove();
      }, 500);
    }, 5000);
  });

  var navLinks = document.querySelectorAll('.nav-link');
  var currentPath = window.location.pathname;
  navLinks.forEach(function (link) {
    var href = link.getAttribute('href');
    if (href && currentPath.startsWith(href) && href !== '/') {
      link.style.background = 'rgba(255, 255, 255, 0.2)';
      link.style.color = 'white';
    }
  });
});
