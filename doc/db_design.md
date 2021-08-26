# Database Design Scheme of GoChat

## 1. GOCHAT_ADMINS

### 1.1. Overview

|  |  |
| :-: | :-: |
| **Table Name** | ADMINS_TB |
| **Primary Key** | UID |
| **Table Constraints** | / |

### 1.2. Detail

| Key | Type | Description| Column Constraints | Optional Value |
| :-: | :-: | :-: | :-: | :-: |
| UID | `VARCHAR(32)` | user's id | PRIMARY KEY | / |
| PWD | `TEXT` | user's password | NOT NULL | / |
| TIME | `TEXT` | timestamp of creating the user | DEFAULT CURRENT_TIMESTAMP | / |

```sql
CREATE TABLE ADMINS_TB (
    UID  VARCHAR(32) PRIMARY KEY,
    PWD  TEXT        NOT NULL,
    TIME TEXT        DEFAULT CURRENT_TIMESTAMP
)
```

## 2. GOCHAT_ROOMS

### 2.1. Overview

|  |  |
| :-: | :-: |
| **Table Name** | ROOMS_TB |
| **Primary Key** | RID |
| **Table Constraints** | / |

### 2.2. Detail

| Key | Type | Description| Column Constraints | Optional Value |
| :-: | :-: | :-: | :-: | :-: |
| RID | `VARCHAR(64)` | id of room | PRIMARY KEY | / |
| ADMIN | `VARCHAR(32)` | administrator's id | REFERENCES ADMINS_TB(UID) | / |
| TOKEN | `VARCHAR(20)` | token of room | NOT NULL | / |
| TIME | `TEXT` | timestamp of creating the room | DEFAULT CURRENT_TIMESTAMP | / |

```sql
CREATE TABLE ROOMS_TB (
    RID   VARCHAR(64) PRIMARY KEY,
    ADMIN VARCHAR(32) REFERENCES ADMINS_TB(UID),
    TOKEN VARCHAR(20) NOT NULL,
    TIME  TEXT        DEFAULT CURRENT_TIMESTAMP
)
```
