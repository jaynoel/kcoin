package oracle_test

import (
	"math/big"
	"strings"
	"testing"

	"crypto/ecdsa"

	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind/backends"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/oracle/testfiles"
	"github.com/kowala-tech/kcoin/client/contracts/bindings/utils"
	"github.com/kowala-tech/kcoin/client/core"
	"github.com/kowala-tech/kcoin/client/crypto"
	"github.com/kowala-tech/kcoin/client/params"
	"github.com/stretchr/testify/suite"
)

var (
	owner, _       = ecdsa.GenerateKey(crypto.S256(), strings.NewReader("ownerAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")) // Needed deterministic in order to link NameHash lib in contract.
	user, _        = ecdsa.GenerateKey(crypto.S256(), strings.NewReader("userAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"))
	initialBalance = new(big.Int).Mul(new(big.Int).SetUint64(100), new(big.Int).SetUint64(params.Kcoin))
)

type OracleMgrSuite struct {
	suite.Suite
	backend          *backends.SimulatedBackend
	oracleMgr        *testfiles.OracleMgr
	resolverMockAddr common.Address
}

func TestOracleMgrSuite(t *testing.T) {
	suite.Run(t, new(OracleMgrSuite))
}

func (suite *OracleMgrSuite) BeforeTest(suiteName, testName string) {
	req := suite.Require()

	backend := backends.NewSimulatedBackend(core.GenesisAlloc{
		crypto.PubkeyToAddress(owner.PublicKey): core.GenesisAccount{
			Balance: initialBalance,
		},
		crypto.PubkeyToAddress(user.PublicKey): core.GenesisAccount{
			Balance: initialBalance,
		},
	})
	req.NotNil(backend)
	suite.backend = backend

	if strings.Contains(testName, "TestDeploy") {
		return
	}

	maxNumOracles := 50
	switch {
	case strings.Contains(testName, "_NotSuperNode"):
		fallthrough
	case strings.Contains(testName, "_Full"):
		maxNumOracles = 1
	}

	transactOpts := bind.NewKeyedTransactor(owner)

	suite.deployStringsLibrary(transactOpts)
	suite.deployNameHashLibrary(transactOpts)
	resolverMockAddress := suite.deployResolverMock(transactOpts)

	// deploy oracle mgr contract
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(50)
	_, _, oracleMgrContract, err := testfiles.DeployOracleMgr(transactOpts, suite.backend, big.NewInt(int64(maxNumOracles)), syncFreq, updatePeriod, resolverMockAddress)
	req.NoError(err)
	req.NotNil(oracleMgrContract)
	suite.oracleMgr = oracleMgrContract

	suite.backend.Commit()
}

func (suite *OracleMgrSuite) TestDeploy() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	suite.deployStringsLibrary(transactOpts)
	suite.deployNameHashLibrary(transactOpts)
	suite.deployResolverMock(transactOpts)

	//deploy oracle mgr contract
	maxNumOracles := big.NewInt(50)
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(50)
	_, _, oracleMgrContract, err := testfiles.DeployOracleMgr(transactOpts, suite.backend, maxNumOracles, syncFreq, updatePeriod, suite.resolverMockAddr)
	req.NoError(err)
	req.NotNil(oracleMgrContract)

	suite.backend.Commit()

	storedMaxNumOracles, err := oracleMgrContract.MaxNumOracles(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedMaxNumOracles)
	req.Equal(maxNumOracles, storedMaxNumOracles)

	storedSyncFreq, err := oracleMgrContract.SyncFrequency(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedSyncFreq)
	req.Equal(syncFreq, storedSyncFreq)

	storedUpdatePeriod, err := oracleMgrContract.UpdatePeriod(&bind.CallOpts{})
	req.NoError(err)
	req.NotNil(storedUpdatePeriod)
	req.Equal(updatePeriod, storedUpdatePeriod)
}

func (suite *OracleMgrSuite) deployResolverMock(transactOpts *bind.TransactOpts) common.Address {
	req := suite.Require()

	mockAddr, _, _, err := testfiles.DeployDomainResolverMock(transactOpts, suite.backend)
	req.NoError(err)
	req.NotZero(mockAddr)
	suite.backend.Commit()

	suite.resolverMockAddr = mockAddr

	return mockAddr
}

func (suite *OracleMgrSuite) deployNameHashLibrary(transactOpts *bind.TransactOpts) common.Address {
	req := suite.Require()

	nameHashLib, _, _, err := utils.DeployNameHash(transactOpts, suite.backend)
	req.NoError(err)
	req.NotZero(nameHashLib)
	suite.backend.Commit()

	return nameHashLib
}

func (suite *OracleMgrSuite) deployStringsLibrary(transactOpts *bind.TransactOpts) common.Address {
	req := suite.Require()

	stringsLibAddr, _, _, err := utils.DeployStrings(transactOpts, suite.backend)
	req.NoError(err)
	req.NotZero(stringsLibAddr)
	suite.backend.Commit()

	return stringsLibAddr
}

func (suite *OracleMgrSuite) TestDeploy_MaxNumOraclesEqualZero() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy oracle mgr contract
	maxNumOracles := common.Big0
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(50)
	_, _, _, err := testfiles.DeployOracleMgr(transactOpts, suite.backend, maxNumOracles, syncFreq, updatePeriod, suite.resolverMockAddr)
	req.Error(err, "max number of oracles must be greater than 0")
}

func (suite *OracleMgrSuite) TestDeploy_SyncFreqGreaterZero_UpdatePeriodZero() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy oracle mgr contract
	maxNumOracles := big.NewInt(50)
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(0)
	_, _, _, err := testfiles.DeployOracleMgr(transactOpts, suite.backend, maxNumOracles, syncFreq, updatePeriod, suite.resolverMockAddr)
	req.Error(err, "update period must be greater than 0 when sync is enabled")
}

func (suite *OracleMgrSuite) TestDeploy_SyncFreqGreaterZero_UpdatePeriodGreaterSyncFreq() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// deploy oracle mgr contract
	maxNumOracles := big.NewInt(50)
	syncFreq := big.NewInt(900)
	updatePeriod := big.NewInt(1000)
	_, _, _, err := testfiles.DeployOracleMgr(transactOpts, suite.backend, maxNumOracles, syncFreq, updatePeriod, suite.resolverMockAddr)
	req.Error(err, "update period must be less or equal than sync freq")
}

func (suite *OracleMgrSuite) TestRegisterOracle_WhenPaused() {
	req := suite.Require()

	transactOpts := bind.NewKeyedTransactor(owner)

	// pause service
	suite.oracleMgr.Pause(transactOpts)

	suite.backend.Commit()

	// register oracle must fail
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.Error(err, "service is paused")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotPaused_Duplicate() {
	req := suite.Require()

	// register an oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// register the same oracle again
	_, err = suite.oracleMgr.RegisterOracle(registerOpts)
	req.Error(err, "duplicate registration")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotPaused_NewCandidate_NotSuperNode() {
	req := suite.Require()

	// register an oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.Error(err, "user is not a super node")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotPaused_NewCandidate_SuperNode_Full() {
	req := suite.Require()

	// register a new oracle to match the max number of oracles
	registerOpts := bind.NewKeyedTransactor(owner)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)

	// register a new oracle
	registerOpts = bind.NewKeyedTransactor(user)
	_, err = suite.oracleMgr.RegisterOracle(registerOpts)
	req.Error(err, "no positions available")
}

func (suite *OracleMgrSuite) TestRegisterOracle_NotPaused_NewCandidate_SuperNode_NotFull() {
	req := suite.Require()

	// register a new oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)
}

func (suite *OracleMgrSuite) TestDeregisterOracle_WhenPaused() {
	req := suite.Require()

	// register oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// pause service
	pauseOpts := bind.NewKeyedTransactor(owner)
	suite.oracleMgr.Pause(pauseOpts)

	suite.backend.Commit()

	// deregister oracle
	deregisterOpts := bind.NewKeyedTransactor(user)
	_, err = suite.oracleMgr.DeregisterOracle(deregisterOpts)
	req.Error(err, "service is paused")
}

func (suite *OracleMgrSuite) TestDeregisterOracle_NotPaused_NotOracle() {
	req := suite.Require()

	// deregister oracle
	deregisterOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.DeregisterOracle(deregisterOpts)
	req.Error(err, "the user is not an oracle")
}

func (suite *OracleMgrSuite) TestDeregisterOracle_NotPaused_Oracle() {
	req := suite.Require()

	// register oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// deregister oracle
	deregisterOpts := bind.NewKeyedTransactor(user)
	_, err = suite.oracleMgr.DeregisterOracle(deregisterOpts)
	req.NoError(err)

	suite.backend.Commit()

	// oracle count must be zero
	count, err := suite.oracleMgr.GetOracleCount(&bind.CallOpts{})
	req.NoError(err)
	req.Zero(count.Uint64())
}

func (suite *OracleMgrSuite) TestSubmitPrice_WhenPaused() {
	req := suite.Require()

	// register oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// pause service
	pauseOpts := bind.NewKeyedTransactor(owner)
	suite.oracleMgr.Pause(pauseOpts)

	suite.backend.Commit()

	priceOpts := bind.NewKeyedTransactor(user)
	_, err = suite.oracleMgr.SubmitPrice(priceOpts, common.Big2)
	req.Error(err, "service is paused")
}

func (suite *OracleMgrSuite) TestSubmitPrice_NotPaused_NotOracle() {
	req := suite.Require()

	priceOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.SubmitPrice(priceOpts, common.Big2)
	req.Error(err, "the user is not an oracle")
}

func (suite *OracleMgrSuite) TestSubmitPrice_NotPaused_Oracle() {
	req := suite.Require()

	// register oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// submit price
	priceOpts := bind.NewKeyedTransactor(user)
	_, err = suite.oracleMgr.SubmitPrice(priceOpts, common.Big2)
	req.NoError(err)
}

func (suite *OracleMgrSuite) TestSubmitPrice_NotPaused_Oracle_SubmitPriceTwiceSameRound() {
	req := suite.Require()

	// register oracle
	registerOpts := bind.NewKeyedTransactor(user)
	_, err := suite.oracleMgr.RegisterOracle(registerOpts)
	req.NoError(err)

	suite.backend.Commit()

	// submit price
	priceOpts := bind.NewKeyedTransactor(user)
	_, err = suite.oracleMgr.SubmitPrice(priceOpts, common.Big2)
	req.NoError(err)

	suite.backend.Commit()

	// submit price
	_, err = suite.oracleMgr.SubmitPrice(priceOpts, common.Big2)
	req.Error(err, "cannot submit a price twice in the same round")
}
