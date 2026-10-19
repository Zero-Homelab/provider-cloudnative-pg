package provider

import (
	"testing"

	cnpgv1 "github.com/cloudnative-pg/cloudnative-pg/api/v1"
)

func TestServiceAccountTokenOption(t *testing.T) {
	t.Run("is optional", func(t *testing.T) {
		if err := syncServiceAccountToken(nil, nil); err != nil {
			t.Fatalf("optional ServiceAccount token returned an error: %v", err)
		}
	})

	t.Run("applies only the intended ServiceAccount field", func(t *testing.T) {
		object := serviceAccountTokenApplyObject("zerotools", "zerotools", true)
		if got := object.Object["automountServiceAccountToken"]; got != true {
			t.Fatalf("automountServiceAccountToken = %v, want true", got)
		}
		metadata, ok := object.Object["metadata"].(map[string]any)
		if !ok || metadata["name"] != "zerotools" || metadata["namespace"] != "zerotools" {
			t.Fatalf("unexpected ServiceAccount identity: %#v", object.Object["metadata"])
		}
		if len(object.Object) != 4 {
			t.Fatalf("ServiceAccount apply object has unexpected fields: %#v", object.Object)
		}
	})
}

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
