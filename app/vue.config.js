const fs = require('fs');

module.exports = {
    devServer: {
        https: {
            cert: fs.readFileSync("../server/server.crt"),
            key: fs.readFileSync("../server/server.key")
        },
        port: 3000
    },
};
