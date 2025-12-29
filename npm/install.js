#!/usr/bin/env node

const fs = require("fs");
const path = require("path");
const os = require("os");

const PLATFORM_MAP = {
  darwin: {
    arm64: "@nikero/updep-darwin-arm64",
    x64: "@nikero/updep-darwin-x64",
  },
  linux: {
    arm64: "@nikero/updep-linux-arm64",
    x64: "@nikero/updep-linux-x64",
  },
  win32: {
    arm64: "@nikero/updep-win32-arm64",
    x64: "@nikero/updep-win32-x64",
  },
};

function getPlatformPackage() {
  const platform = os.platform();
  const arch = os.arch();

  if (!(platform in PLATFORM_MAP)) {
    throw new Error(
      `Unsupported platform: ${platform}. updep supports ${Object.keys(PLATFORM_MAP).join(", ")}.`
    );
  }
  if (!PLATFORM_MAP[platform][arch]) {
    throw new Error(
      `Unsupported architecture: ${arch} on ${platform}. updep supports ${Object.keys(PLATFORM_MAP[platform]).join(", ")}.`
    );
  }

  return PLATFORM_MAP[platform][arch];
}

function install() {
  try {
    const packageName = getPlatformPackage();
    const binName = `updep${os.platform() === "win32" ? ".exe" : ""}`;

    let binaryPath;
    try {
      binaryPath = require.resolve(`${packageName}/bin/${binName}`);
    } catch (e) {
      console.warn(
        `Warning: Could not find platform-specific package ${packageName}.`
      );
      console.warn(
        "This is expected during development. In production, npm will install the correct package."
      );
      return;
    }

    // Copy binary to bin directory
    const targetPath = path.join(__dirname, "bin", binName);
    fs.copyFileSync(binaryPath, targetPath);
    fs.chmodSync(targetPath, 0o755);

    console.log(
      `✓ updep installed successfully for ${os.platform()}-${os.arch()}`
    );
  } catch (error) {
    console.error("Error installing updep:", error.message);
    process.exit(1);
  }
}

install();
