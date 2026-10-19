package provider

import (
	"testing"

	cnpgv1 "github.com/cloudnative-pg/cloudnative-pg/api/v1"
)

func TestApplySuperuserAccess(t *testing.T) {
	t.Run("preserves CNPG default when omitted", func(t *testing.T) {
		pg := &cnpgv1.Cluster{}
		applySuperuserAccess(pg, nil)
		if pg.Spec.EnableSuperuserAccess != nil {
			t.Fatal("enableSuperuserAccess changed when parameter was omitted")
		}
	})

	t.Run("enables CNPG-managed postgres credentials", func(t *testing.T) {
		pg := &cnpgv1.Cluster{}
		enabled := true
		applySuperuserAccess(pg, &enabled)
		if pg.Spec.EnableSuperuserAccess == nil || !*pg.Spec.EnableSuperuserAccess {
			t.Fatal("enableSuperuserAccess is not true")
		}
		if pg.Spec.SuperuserSecret != nil {
			t.Fatal("provider must leave Secret generation to CNPG")
		}
	})
}
