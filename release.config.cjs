/**
 * @type {import('semantic-release').GlobalConfig}
 */
module.exports = {
    branches: ["master"],
    plugins: [
        '@semantic-release/commit-analyzer',
        '@semantic-release/release-notes-generator',
        '@semantic-release/git', {
            "message": "chore(release): ${nextRelease.version} [skip ci]\n\n${nextRelease.notes}"
        },
        '@semantic-release/github', {
            "assets": [
                {"path": "release/wordler-windows-amd64.exe", "label": "Windows 64bit"},
                {"path": "release/wordler-darwin-amd64", "label": "MacOS 64bit Intel"},
                {"path": "release/wordler-darwin-arm64", "label": "MacOS 64bit Apple Silion"},
                {"path": "release/wordler-linux-amd64", "label": "Linux 64bit"}
            ]
        }
    ]
};
