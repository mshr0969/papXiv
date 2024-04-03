
-- +migrate Up
CREATE TABLE paper_subjects (
    paper_id CHAR(36) NOT NULL,
    subject_id INT NOT NULL,
    FOREIGN KEY (paper_id) REFERENCES papers(id) ON DELETE CASCADE,
    FOREIGN KEY (subject_id) REFERENCES subjects(id) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS paper_subjects;
