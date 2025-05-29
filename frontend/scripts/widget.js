document.addEventListener('DOMContentLoaded', () => {
  const ratesGrid = document.getElementById('ratesGrid');
  const lastUpdateElem = document.getElementById('lastUpdate');
  const workingHoursElem = document.getElementById('workingHours');
  const telegramButton = document.getElementById('telegramButton');

  fetch('http://localhost:8080/latest')
    .then(response => response.json())
    .then(data => {
      lastUpdateElem.textContent = data.date || '--';
      workingHoursElem.textContent = data.hours + ' МСК' || '--';

      data.rates.forEach(rate => {
        const card = document.createElement('div');
        card.className = 'rate-card';

        const range = document.createElement('div');
        range.className = 'range';
        if (rate.max === -1) {
          range.textContent = `от ${rate.min}¥`;
        } else {
          range.textContent = `${rate.min}¥ - ${rate.max}¥`;
        }

        const rateValue = document.createElement('div');
        rateValue.className = 'rate-value';
        rateValue.textContent = `${rate.rate} ₽`;

        card.appendChild(range);
        card.appendChild(rateValue);
        ratesGrid.appendChild(card);
      });
    })
    .catch(error => {
      console.error('Ошибка при получении данных:', error);
      // Фолбэк-данные
      const fallbackRates = [
        { min: 200, max: 999, rate: 12.00 },
        { min: 1000, max: 2999, rate: 11.85 },
        { min: 3000, max: -1, rate: 11.75 }
      ];
      lastUpdateElem.textContent = new Date().toLocaleDateString('ru-RU');
      workingHoursElem.textContent = '13:00 - 20:00 (МСК)';

      fallbackRates.forEach(rate => {
        const card = document.createElement('div');
        card.className = 'rate-card';

        const range = document.createElement('div');
        range.className = 'range';
        if (rate.max === -1) {
          range.textContent = `от ${rate.min}¥`;
        } else {
          range.textContent = `${rate.min}¥ - ${rate.max}¥`;
        }

        const rateValue = document.createElement('div');
        rateValue.className = 'rate-value';
        rateValue.textContent = `${rate.rate} ₽`;

        card.appendChild(range);
        card.appendChild(rateValue);
        ratesGrid.appendChild(card);
      });
    });

  telegramButton.addEventListener('click', () => {
    window.open('https://t.me/+-tkxgvzLYT42Yzgy', '_blank');
  });
});
