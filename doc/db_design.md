# Database Design Scheme of GoChat

## 1. GOCHAT_ADMINS

### 1.1. Overview

|  |  |
| :-: | :-: |
| **Table Name** | GOCHAT_ADMINS |
| **Primary Key** | / |
| **Table Constraints** | / |

### 1.2. Detail

| Key | Type | Description| Column Constraints | Optional Value |
| :-: | :-: | :-: | :-: | :-: |
| UID | `VARCHAR(32)` | user's id | NOT NULL UNIQUE | / |
| PWD | `VARCHAR(20)` | user's password | NOT NULL | / |
| TIME | `TIMESTAMP(3)` | timestamp of creating the user | DEFAULT NOW() | / |

```sql
CREATE TABLE GOCHAT_ADMINS (
    UID VARCHAR(32) NOT NULL UNIQUE,
    PWD VARCHAR(20) NOT NULL,
    TIME TIMESTAMP(3) DEFAULT NOW()
)
```

## 2. GOCHAT_ROOMS

### 2.1. Overview

|  |  |
| :-: | :-: |
| **Table Name** | GOCHAT_ROOMS |
| **Primary Key** | / |
| **Table Constraints** | / |

### 2.2. Detail

| Key | Type | Description| Column Constraints | Optional Value |
| :-: | :-: | :-: | :-: | :-: |
| RID | `VARCHAR(64)` | id of room | NOT NULL UNIQUE | / |
| ADMIN | `VARCHAR(32)` | administrator's id | REFERENCES GOCHAT_ADMINS (UID) | / |
| TOKEN | `VARCHAR(20)` | token of room | NOT NULL | / |
| TIME | `TIMESTAMP(3)` | timestamp of creating the room | DEFAULT NOW() | / |

```sql
CREATE TABLE GOCHAT_ROOMS (
    RID VARCHAR(64) NOT NULL UNIQUE,
    ADMIN VARCHAR(32) REFERENCES GOCHAT_ADMINS (UID),
    TOKEN VARCHAR(20) NOT NULL,
    TIME TIMESTAMP(3) DEFAULT NOW()
)
```
