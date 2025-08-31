package cache

import "testing"

func TestRedisCache_Has(t *testing.T) {
	err := testRedisCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testRedisCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo found in cache, and it shouldn't be there")
	}

	err = testRedisCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}
	inCache, err = testRedisCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("foo not found in cache, but it should be there")
	}
}

func TestRedisCache_Get(t *testing.T) {
	err := testRedisCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}
	x, err := testRedisCache.Get("foo")
	if err != nil {
		t.Error(err)
	}
	if x != "bar" {
		t.Error("did not get correct value from cache")
	}
}

func TestRedisCache_Forget(t *testing.T) {
	err := testRedisCache.Set("alpha", "beta")
	if err != nil {
		t.Error(err)
	}
	err = testRedisCache.Forget("alpha")
	if err != nil {
		t.Error(err)
	}
	inCache, err := testRedisCache.Has("alpha")
	if inCache {
		t.Error("alpha found in cache, and it should not be there")
	}
}

func TestRedisCache_Empty(t *testing.T) {
	err := testRedisCache.Set("gamma", "delta")
	if err != nil {
		t.Error(err)
	}
	err = testRedisCache.Empty()
	if err != nil {
		t.Error(err)
	}

	inCache, err := testRedisCache.Has("gamma")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("gamma found in cache, but it should not be there")
	}
}

func TestRedisCache_EmptyByMatch(t *testing.T) {
	err := testRedisCache.Set("epsilon", "zeta")
	if err != nil {
		t.Error(err)
	}
	err = testRedisCache.Set("eta", "theta")
	if err != nil {
		t.Error(err)
	}

	err = testRedisCache.Set("iota", "kappa")
	if err != nil {
		t.Error(err)
	}

	err = testRedisCache.EmptyByMatch("e")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testRedisCache.Has("epsilon")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("epsilon found in cache, but it should not be there")
	}

	inCache, err = testRedisCache.Has("eta")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("eta found in cache, but it should not be there")
	}

	inCache, err = testRedisCache.Has("iota")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("iota not found in cache, but it should be there")
	}
}

func TestEncodeDecode(t *testing.T) {
	entry := Entry{}
	entry["foo"] = "bar"
	bytes, err := encode(entry)
	if err != nil {
		t.Error(err)
	}
	_, err = decode(string(bytes))
	if err != nil {
		t.Error(err)
	}
}
