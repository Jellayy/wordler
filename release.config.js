const config = {
    branches: ['master'],
    plugins: [
        '@semantic-release/commit-analyzer',
        '@semantic-release/release-notes-generator',
        '@semantic-release/git', {
            "message": "chore(release): ${nextRelease.version} [skip ci]\n\n${nextRelease.notes}"
        },
        '@semantic-release/github', {
            "assets": [
                {
                    "path": "releases/*",
                    "label": "Release ${nextRelease.version}"
                }
            ]
        }
    ]
};

module.exports = config;
