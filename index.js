yaml = require('js-yaml');
fs   = require('fs');

// Get document, or throw exception on error
try {
  var doc = yaml.safeLoad(fs.readFileSync('./exercises.yaml', 'utf8'));
  console.log(doc);
} catch (e) {
  console.log(e);
}
