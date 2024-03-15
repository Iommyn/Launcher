SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";

--
-- База данных: `corecraft`
--

--
-- Структура таблицы `pending_transactions`
--

CREATE TABLE `pending_transactions` (
  `id` int NOT NULL,
  `balance_id` int NOT NULL,
  `user_id` int NOT NULL,
  `amount` bigint NOT NULL,
  `status` varchar(7) COLLATE utf8mb4_general_ci NOT NULL,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `servers`
--

CREATE TABLE `servers` (
  `id` int NOT NULL,
  `category_id` int NOT NULL,
  `server_name` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
  `mini_desc` text COLLATE utf8mb4_general_ci NOT NULL,
  `description` text COLLATE utf8mb4_general_ci NOT NULL,
  `list_mods` text COLLATE utf8mb4_general_ci NOT NULL,
  `link` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
  `isPvP` tinyint(1) NOT NULL,
  `worlds` text COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_wipe` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `servers_category`
--

CREATE TABLE `servers_category` (
  `id` int NOT NULL,
  `version` varchar(8) COLLATE utf8mb4_general_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `servers_data`
--

CREATE TABLE `servers_data` (
  `id` int NOT NULL,
  `category_id` int NOT NULL,
  `server_ip` int NOT NULL,
  `server_port` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `servers_image`
--

CREATE TABLE `servers_image` (
  `id` int NOT NULL,
  `category_id` int NOT NULL,
  `image_link` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
  `image_one` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
  `image_two` varchar(256) COLLATE utf8mb4_general_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `servers_rcon`
--

CREATE TABLE `servers_rcon` (
  `id` int NOT NULL,
  `category_id` int NOT NULL,
  `rcon_ip` varchar(15) COLLATE utf8mb4_general_ci NOT NULL,
  `rcon_port` int NOT NULL,
  `rcon_password` varchar(256) COLLATE utf8mb4_general_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `users`
--

CREATE TABLE `users` (
  `id` int NOT NULL,
  `email` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `username` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
  `isAdmin` tinyint(1) NOT NULL DEFAULT '0',
  `ip` varchar(15) COLLATE utf8mb4_general_ci NOT NULL,
  `uuid` char(36) COLLATE utf8mb4_general_ci NOT NULL,
  `accessToken` varchar(32) COLLATE utf8mb4_general_ci NOT NULL,
  `serverID` varchar(41) COLLATE utf8mb4_general_ci NOT NULL,
  `reg_data` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_login_site` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Дамп данных таблицы `users`

-- --------------------------------------------------------
--
-- Структура таблицы `user_2fa`
--

CREATE TABLE `user_2fa` (
  `id` int NOT NULL,
  `user_id` int NOT NULL,
  `2fa` tinyint(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `user_economy`
--

CREATE TABLE `user_economy` (
  `id` int NOT NULL,
  `user_id` int NOT NULL,
  `balance` int NOT NULL DEFAULT '0',
  `isFreeze` tinyint NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `user_sessions`
--

CREATE TABLE `user_sessions` (
  `id` char(36) COLLATE utf8mb4_general_ci NOT NULL,
  `username` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `refresh_token` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `user_agent` text COLLATE utf8mb4_general_ci NOT NULL,
  `client_ip` varchar(15) COLLATE utf8mb4_general_ci NOT NULL,
  `isBlocked` tinyint NOT NULL DEFAULT '0',
  `expires_at` datetime NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


-- --------------------------------------------------------

--
-- Структура таблицы `user_vote`
--

CREATE TABLE `user_vote` (
  `id` int NOT NULL,
  `user_id` int NOT NULL,
  `votes_count` int NOT NULL,
  `last_data_votes` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `pending_transactions`
--
ALTER TABLE `pending_transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `balance_id` (`balance_id`);

--
-- Индексы таблицы `servers`
--
ALTER TABLE `servers`
  ADD PRIMARY KEY (`id`),
  ADD KEY `servers_fk0` (`category_id`);

--
-- Индексы таблицы `servers_category`
--
ALTER TABLE `servers_category`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `servers_data`
--
ALTER TABLE `servers_data`
  ADD PRIMARY KEY (`id`),
  ADD KEY `servers_data_fk0` (`category_id`);

--
-- Индексы таблицы `servers_image`
--
ALTER TABLE `servers_image`
  ADD PRIMARY KEY (`id`),
  ADD KEY `servers_image_fk0` (`category_id`);

--
-- Индексы таблицы `servers_rcon`
--
ALTER TABLE `servers_rcon`
  ADD PRIMARY KEY (`id`),
  ADD KEY `servers_rcon_fk0` (`category_id`);

--
-- Индексы таблицы `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `user_2fa`
--
ALTER TABLE `user_2fa`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_2fa_fk0` (`user_id`);

--
-- Индексы таблицы `user_economy`
--
ALTER TABLE `user_economy`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_economy_fk0` (`user_id`);

--
-- Индексы таблицы `user_sessions`
--
ALTER TABLE `user_sessions`
  ADD KEY `username` (`username`);

--
-- Индексы таблицы `user_vote`
--
ALTER TABLE `user_vote`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_vote_fk0` (`user_id`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `pending_transactions`
--
ALTER TABLE `pending_transactions`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT для таблицы `servers`
--
ALTER TABLE `servers`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT для таблицы `servers_category`
--
ALTER TABLE `servers_category`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT для таблицы `servers_data`
--
ALTER TABLE `servers_data`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT для таблицы `servers_image`
--
ALTER TABLE `servers_image`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT для таблицы `servers_rcon`
--
ALTER TABLE `servers_rcon`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;


ALTER TABLE `users`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4441;

--
-- AUTO_INCREMENT для таблицы `user_economy`
--
ALTER TABLE `user_economy`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- Ограничения внешнего ключа таблицы `pending_transactions`
--
ALTER TABLE `pending_transactions`
  ADD CONSTRAINT `pending_transactions_ibfk_1` FOREIGN KEY (`balance_id`) REFERENCES `user_economy` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT;

--
-- Ограничения внешнего ключа таблицы `user_economy`
--
ALTER TABLE `user_economy`
  ADD CONSTRAINT `user_economy_fk0` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
  
--
-- Ограничения внешнего ключа таблицы `user_sessions`
--
ALTER TABLE `user_sessions`
  ADD CONSTRAINT `user_sessions_ibfk_1` FOREIGN KEY (`username`) REFERENCES `users` (`username`) ON DELETE RESTRICT ON UPDATE RESTRICT;

--
-- Ограничения внешнего ключа таблицы `user_vote`
--
ALTER TABLE `user_vote`
  ADD CONSTRAINT `user_vote_fk0` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
