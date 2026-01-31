
CREATE TABLE administrator
(
  id uuid NOT NULL PRIMARY KEY
);

CREATE TABLE club
(
  id   uuid NOT NULL,
  name TEXT NOT NULL,
  slug TEXT NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE club_organiser
(
  club_id   uuid NOT NULL,
  runner_id uuid NOT NULL,
  PRIMARY KEY (club_id, runner_id)
);

CREATE TABLE club_runner
(
  club_id   uuid NOT NULL,
  runner_id uuid NOT NULL,
  PRIMARY KEY (club_id, runner_id)
);

CREATE TABLE event
(
  club_id uuid NOT NULL,
  id      uuid NOT NULL,
  name    TEXT NOT NULL,
  slug    TEXT NOT NULL,
  date    DATE NOT NULL,
  time    TIME NULL    ,
  PRIMARY KEY (id)
);

CREATE TABLE event_finisher
(
  id               uuid   NOT NULL PRIMARY KEY,
  event_id         uuid   NOT NULL,
  position         BIGINT NOT NULL,
  time_seconds     BIGINT NOT NULL,
  event_starter_id uuid   NULL    
);

CREATE TABLE event_starter
(
  id                     uuid   NOT NULL,
  event_id               uuid   NOT NULL,
  projected_time_seconds BIGINT NOT NULL,
  runner_id              uuid   NULL    ,
  placeholder_name       TEXT   NULL    ,
  PRIMARY KEY (id)
);

CREATE TABLE login
(
  id         uuid NOT NULL,
  provider   TEXT NOT NULL,
  identifier TEXT NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE runner
(
  id          uuid NOT NULL,
  given_name  TEXT NULL    ,
  family_name TEXT NULL    ,
  PRIMARY KEY (id)
);

CREATE TABLE runner_login
(
  login_id  uuid NOT NULL,
  runner_id uuid NOT NULL,
  PRIMARY KEY (login_id, runner_id)
);

CREATE TABLE runner_time
(
  runner_id    uuid   NOT NULL,
  id           uuid   NOT NULL,
  date         DATE   NOT NULL,
  time         TIME   NULL    ,
  time_seconds BIGINT NOT NULL,
  name         TEXT   NOT NULL,
  description  TEXT   NULL    ,
  PRIMARY KEY (id)
);

ALTER TABLE runner_login
  ADD CONSTRAINT FK_login_TO_runner_login
    FOREIGN KEY (login_id)
    REFERENCES login (id);

ALTER TABLE runner_login
  ADD CONSTRAINT FK_runner_TO_runner_login
    FOREIGN KEY (runner_id)
    REFERENCES runner (id);

ALTER TABLE administrator
  ADD CONSTRAINT FK_login_TO_administrator
    FOREIGN KEY (id)
    REFERENCES login (id);

ALTER TABLE event
  ADD CONSTRAINT FK_club_TO_event
    FOREIGN KEY (club_id)
    REFERENCES club (id);

ALTER TABLE club_organiser
  ADD CONSTRAINT FK_club_TO_club_organiser
    FOREIGN KEY (club_id)
    REFERENCES club (id);

ALTER TABLE club_organiser
  ADD CONSTRAINT FK_runner_TO_club_organiser
    FOREIGN KEY (runner_id)
    REFERENCES runner (id);

ALTER TABLE event_starter
  ADD CONSTRAINT FK_runner_TO_event_starter
    FOREIGN KEY (runner_id)
    REFERENCES runner (id);

ALTER TABLE event_starter
  ADD CONSTRAINT FK_event_TO_event_starter
    FOREIGN KEY (event_id)
    REFERENCES event (id);

ALTER TABLE club_runner
  ADD CONSTRAINT FK_club_TO_club_runner
    FOREIGN KEY (club_id)
    REFERENCES club (id);

ALTER TABLE club_runner
  ADD CONSTRAINT FK_runner_TO_club_runner
    FOREIGN KEY (runner_id)
    REFERENCES runner (id);

ALTER TABLE event_finisher
  ADD CONSTRAINT FK_event_TO_event_finisher
    FOREIGN KEY (event_id)
    REFERENCES event (id);

ALTER TABLE event_finisher
  ADD CONSTRAINT FK_event_starter_TO_event_finisher
    FOREIGN KEY (event_starter_id)
    REFERENCES event_starter (id);

ALTER TABLE runner_time
  ADD CONSTRAINT FK_runner_TO_runner_time
    FOREIGN KEY (runner_id)
    REFERENCES runner (id);
