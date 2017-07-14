const sighting = {
  id: 1,
  location: "Roswell, New Mexico"
};

export default function fetch(url) {
  return new Promise((resolve, reject) => {
    switch (url) {
      case "/api/reports/1/":
        resolve(sighting);
      default:
        reject("Invalid url");
    }
  });
}
