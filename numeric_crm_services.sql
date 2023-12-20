-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Dec 07, 2023 at 01:13 AM
-- Server version: 8.2.0
-- PHP Version: 8.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `numeric_crm_services`
--

-- --------------------------------------------------------

--
-- Table structure for table `actors`
--

CREATE TABLE `actors` (
  `id` bigint UNSIGNED NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role_id` int UNSIGNED DEFAULT '2',
  `verified` enum('true','false') DEFAULT 'false',
  `active` enum('true','false') DEFAULT 'false',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `actors`
--

INSERT INTO `actors` (`id`, `username`, `password`, `role_id`, `verified`, `active`, `created_at`, `updated_at`) VALUES
(1, 'superadmin', '$2a$12$26EAzIyDQ5YZKE6pAeHKoeugjQ98lkXNPOEFjJikPmQReEMK4snoW', 1, 'true', 'true', '2023-06-02 04:41:18', '2023-06-04 05:18:58'),
(3, 'be', '$2a$12$LLGqut.2ghWEIq7KbqmBgexbC6tys1S52GE7dqghRx/Bn8cvaen2u', 2, 'true', 'true', '2023-06-02 07:33:07', '2023-06-03 09:39:09'),
(4, 'se', '$2a$12$3Cmx2UK/HVCp1kg0z2yLL.V3DcMN/BoBN4jYhjf1I4J3gOHlARDe6', 2, 'false', 'false', '2023-06-02 11:07:55', '2023-06-03 09:17:29'),
(6, 's', '$2a$12$jHCAyoN6vT9yyDSzagl6BOl0MbgdICn1fz1GqW3dRlNuhZvegjDvO', 2, 'false', 'false', '2023-06-02 11:12:35', '2023-06-02 11:12:35'),
(9, 'd', '$2a$12$o5HvC9FBOTQ7Yc/Eyac2HOR36JGnWUMIwipK2YfV6pG25UVK2/v.e', 2, 'false', 'false', '2023-06-04 07:16:11', '2023-06-04 07:16:11'),
(17, 'hi', 'hai', 2, 'false', 'false', '2023-10-06 02:47:48', '2023-10-06 02:47:48'),
(18, 'h0', 'hai', 2, 'false', 'false', '2023-10-06 02:49:19', '2023-10-06 02:49:19'),
(19, 'hai', 'hai', 2, 'false', 'false', '2023-10-06 06:21:35', '2023-10-06 06:21:35'),
(21, 'ence', 'k', 2, 'false', 'false', '2023-10-06 06:25:25', '2023-10-06 06:29:05'),
(25, 'ka', 'ejcnec', 2, 'false', 'false', '2023-10-06 06:28:44', '2023-10-06 06:28:44'),
(28, 'en', 'ejcnec', 2, 'false', 'false', '2023-10-06 06:29:38', '2023-10-06 06:29:54'),
(30, 'j', 'ejcnec', 2, 'false', 'false', '2023-10-06 06:30:05', '2023-10-06 06:30:05'),
(36, 'hahaha', 'k', 2, 'false', 'false', '2023-10-06 06:31:28', '2023-10-06 06:31:28'),
(38, 'hahah', 'k', 2, 'false', 'false', '2023-10-06 06:31:50', '2023-10-06 06:31:50'),
(39, 'w', 'hai', 2, 'true', 'false', '2023-10-06 06:32:25', '2023-10-06 08:22:37'),
(41, 'h', 'h', 2, 'false', 'false', '2023-10-06 18:20:10', '2023-10-06 18:20:10'),
(42, 'ha', '$2a$12$mRjyaHsVCXRpaJ6nKb2uzuqE1c6e/kQG/TtJgB0g8vK7feLrsWxZi', 2, 'false', 'false', '2023-10-22 07:45:27', '2023-10-22 07:45:27'),
(43, 'hjnnnnnn', '$2a$12$XwbeoGQPtTn8YwjY0yF5oeDLhCOuF69KDEAFz8ZbP.u9RYkYBmeE2', 2, 'false', 'false', '2023-10-22 07:58:28', '2023-10-22 07:58:28'),
(44, 'hjnnnnn', '$2a$12$MbzRLn0jyC0ryqsEyIBeNu0YNJXP2RNKA4ux..dBCZOmWAjy5c9ga', 2, 'false', 'false', '2023-10-22 08:00:55', '2023-10-22 08:00:55'),
(45, 'jjj', 'ii', 2, 'false', 'false', '2023-11-08 04:25:01', '2023-11-08 04:25:01');

-- --------------------------------------------------------

--
-- Table structure for table `actor_roles`
--

CREATE TABLE `actor_roles` (
  `id` int UNSIGNED NOT NULL,
  `role_name` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `actor_roles`
--

INSERT INTO `actor_roles` (`id`, `role_name`) VALUES
(2, 'admin'),
(1, 'superadmin');

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `id` bigint UNSIGNED NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  `email` varchar(255) NOT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`id`, `first_name`, `last_name`, `email`, `avatar`, `created_at`, `updated_at`) VALUES
(17, 'Michael', 'Lawson', 'michael.lawson@reqres.in', 'https://reqres.in/img/faces/7-image.jpg', '2023-10-09 02:35:42', '2023-10-09 02:35:42'),
(18, 'Lindsay', 'Ferguson', 'lindsay.ferguson@reqres.in', 'https://reqres.in/img/faces/8-image.jpg', '2023-10-09 02:35:42', '2023-10-09 02:35:42'),
(19, 'Tobias', 'Funke', 'tobias.funke@reqres.in', 'https://reqres.in/img/faces/9-image.jpg', '2023-10-09 02:35:42', '2023-10-09 02:35:42'),
(20, 'Byron', 'Fields', 'byron.fields@reqres.in', 'https://reqres.in/img/faces/10-image.jpg', '2023-10-09 02:35:42', '2023-10-09 02:35:42'),
(21, 'George', 'Edwards', 'george.edwards@reqres.in', 'https://reqres.in/img/faces/11-image.jpg', '2023-10-09 02:35:42', '2023-10-09 02:35:42'),
(22, 'Rachel', 'Howell', 'rachel.howell@reqres.in', 'https://reqres.in/img/faces/12-image.jpg', '2023-10-09 02:35:42', '2023-10-09 02:35:42'),
(23, 'ok', 'hahah', 'hello@ok.k', 'word', '2023-10-09 03:17:53', '2023-10-09 03:17:53');

-- --------------------------------------------------------

--
-- Table structure for table `sessions`
--

CREATE TABLE `sessions` (
  `activity_id` varchar(64) NOT NULL,
  `agent` varchar(255) NOT NULL,
  `issued_at` timestamp NULL DEFAULT NULL,
  `expired_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `sessions`
--

INSERT INTO `sessions` (`activity_id`, `agent`, `issued_at`, `expired_at`) VALUES
('23495761-54de-4b5d-b994-530d4d9c70f1', 'PostmanRuntime/7.35.0', '2023-12-06 11:44:40', '2023-12-06 12:44:41'),
('2a209bfd-ea0e-44e1-bcb4-bdd9bea3eaf4', 'PostmanRuntime/7.35.0', '2023-12-06 11:57:06', '2023-12-06 12:57:07'),
('3e48477e-0ea4-4ee7-9036-3c2707f84ef7', 'PostmanRuntime/7.35.0', '2023-12-06 11:25:56', '2023-12-07 11:25:56'),
('54223001-966d-4a2a-896e-4a11ae6635bd', 'PostmanRuntime/7.35.0', '2023-12-06 13:21:57', '2023-12-07 13:21:57'),
('545ae3c3-28da-4ec6-a38d-fc0364f5a6a6', 'PostmanRuntime/7.35.0', '2023-12-06 13:12:50', '2023-12-07 13:12:50'),
('8959cccd-3725-4ab7-8267-962847cec5b5', 'PostmanRuntime/7.35.0', '2023-12-06 14:55:31', '2023-12-07 14:55:32'),
('8d4e99d8-5b11-4bc8-9edf-528168473b3f', 'PostmanRuntime/7.35.0', '2023-12-06 11:42:15', '2023-12-06 12:42:16'),
('a28910fe-a69c-4328-88c2-43aa74c25772', 'PostmanRuntime/7.35.0', '2023-12-06 11:26:21', '2023-12-07 11:26:21'),
('a4725f50-c94a-4049-8bc2-4c02507345cb', 'PostmanRuntime/7.35.0', '2023-12-06 11:18:30', '2023-12-07 11:18:30'),
('a86d5f35-5578-4230-afb8-4abacc34e922', 'PostmanRuntime/7.35.0', '2023-12-06 14:51:12', '2023-12-07 14:51:13'),
('c921e2c9-e01b-4095-938f-9f060aad6c27', 'PostmanRuntime/7.35.0', '2023-12-06 13:30:46', '2023-12-07 13:30:46'),
('d71ff4b0-312e-4898-ac08-626cb0463525', 'PostmanRuntime/7.35.0', '2023-12-06 13:39:33', '2023-12-07 13:39:34'),
('fb794116-d469-4cdc-b06e-de675e0da24e', 'PostmanRuntime/7.35.0', '2023-12-06 13:11:55', '2023-12-06 14:11:55');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `actors`
--
ALTER TABLE `actors`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD KEY `actor_role` (`role_id`),
  ADD KEY `idx_username_actor` (`username`);

--
-- Indexes for table `actor_roles`
--
ALTER TABLE `actor_roles`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `role_name` (`role_name`);

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD KEY `idx_first_name` (`first_name`),
  ADD KEY `idx_last_name` (`last_name`),
  ADD KEY `customer_index_email` (`email`);

--
-- Indexes for table `sessions`
--
ALTER TABLE `sessions`
  ADD PRIMARY KEY (`activity_id`),
  ADD UNIQUE KEY `activity` (`activity_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `actors`
--
ALTER TABLE `actors`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=46;

--
-- AUTO_INCREMENT for table `actor_roles`
--
ALTER TABLE `actor_roles`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `actors`
--
ALTER TABLE `actors`
  ADD CONSTRAINT `actor_role` FOREIGN KEY (`role_id`) REFERENCES `actor_roles` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
