package i18n

import "fmt"

// Message keys
const (
	MsgConflictFound   = "conflict_found"
	MsgChooseAction    = "choose_action"
	MsgBackupSuffix    = "backup_suffix"
	MsgBackingUp       = "backing_up"
	MsgRemoving        = "removing"
	MsgDeployComplete  = "deploy_complete"
	MsgSnapshotCreated = "snapshot_created"
	MsgSnapshotUpdated = "snapshot_updated"
)

var messages = map[string]map[string]string{
	MsgConflictFound: {
		"zh": "⚠ 目标目录发现非 Cast 管理的同名文件：%s",
		"en": "⚠ Non-Cast-managed file conflicts: %s",
	},
	MsgChooseAction: {
		"zh": "  [1] 备份（默认）\n  [2] 移除\n请选择 [1]: ",
		"en": "  [1] Backup (default)\n  [2] Remove\nChoose [1]: ",
	},
	MsgBackupSuffix: {
		"zh": "备份后缀 [.bak]: ",
		"en": "Backup suffix [.bak]: ",
	},
	MsgBackingUp: {
		"zh": "  📦 备份: %s → %s",
		"en": "  📦 Backup: %s → %s",
	},
	MsgRemoving: {
		"zh": "  🗑  移除: %s",
		"en": "  🗑  Remove: %s",
	},
	MsgDeployComplete: {
		"zh": "  ✅ 部署完成",
		"en": "  ✅ Deploy complete",
	},
	MsgSnapshotCreated: {
		"zh": "  📝 创建 snapshot: .cast/snapshot.json",
		"en": "  📝 Created snapshot: .cast/snapshot.json",
	},
	MsgSnapshotUpdated: {
		"zh": "  📝 更新 snapshot: .cast/snapshot.json",
		"en": "  📝 Updated snapshot: .cast/snapshot.json",
	},
}

// Msg returns a localized message by key and language code.
func Msg(key, lang string) string {
	m, ok := messages[key]
	if !ok {
		return key
	}
	s, ok := m[lang]
	if !ok {
		s = m["en"]
	}
	return s
}

// Msgf returns a formatted localized message.
func Msgf(key, lang string, args ...interface{}) string {
	return fmt.Sprintf(Msg(key, lang), args...)
}
