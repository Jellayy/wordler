/**
 * @type {import('semantic-release').GlobalConfig}
 */
module.exports = {
    branches: ["master"],
    plugins: [
        '@semantic-release/commit-analyzer',
        '@semantic-release/release-notes-generator',
        [
            '@semantic-release/git',
            {
                "message": "chore(release): ${nextRelease.version} [skip ci]\n\n${nextRelease.notes}"
            }
        ],
        [
            '@semantic-release/github',
            {
                "assets": [
                    {"path": "release/wordler-windows-amd64.exe", "label": "wordler-${nextRelease.version}-windows-amd64.exe"},
                    {"path": "release/wordler-darwin-amd64", "label": "wordler-${nextRelease.version}-darwin-amd64"},
                    {"path": "release/wordler-darwin-arm64", "label": "wordler-${nextRelease.version}-darwin-arm64"},
                    {"path": "release/wordler-linux-amd64", "label": "wordler-${nextRelease.version}-linux-amd64"}
                ]
            }
        ]
    ]
};
